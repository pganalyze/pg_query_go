require 'bundler'

def underscore(str)
  str
  .gsub(/([A-Z]+)([A-Z][a-z])/,'\1_\2')
  .gsub(/([a-z\d])([A-Z])/,'\1_\2')
  .tr('-', '_')
  .gsub(/\s/, '_')
  .gsub(/__+/, '_')
  .downcase
end

def classify(str)
  str
  .gsub(/([A-Z]+)/, '_\1')
  .split('_').collect(&:capitalize).join
end

class Generator
  def initialize(pgdir)
    @pgdir = pgdir
  end

  def generate_nodetypes!
    inside = false
    @nodetypes = []

    lines = File.read(File.join(@pgdir, '/src/include/nodes/nodes.h'))
    lines.each_line do |line|
      if inside
        if line[/([A-z_]+)(\s+=\s+\d+)?,/]
          @nodetypes << $1[2..-1] # Without T_ prefix
        elsif line == "} NodeTag;\n"
          inside = false
        end
      elsif line == "typedef enum NodeTag\n"
        inside = true
      end
    end
  end

  IGNORE_LIST = [
    'Node', 'NodeTag', 'varlena', 'IntArray', 'nameData', 'bool',
    'sig_atomic_t', 'size_t', 'varatt_indirect',
    'NullTest', 'BooleanTest', # These fail the compilation process - tbd why
  ]

  def generate_defs!
    @target_struct = nil
    @open_comment = false

    @enum_defs = {}
    @struct_defs = {}
    @typedefs = ''

    ['nodes/parsenodes', 'nodes/plannodes', 'nodes/primnodes', 'nodes/relation',
     'nodes/nodes', 'nodes/params', 'access/attnum', 'c', 'postgres', 'postgres_ext',
     'storage/block', 'access/sdir'].each do |group|
      @target_group = group
      @struct_defs[@target_group] = {}
      @enum_defs[@target_group] = {}

      lines = File.read(File.join(@pgdir, format('/src/include/%s.h', @target_group)))
      lines.each_line do |line|
        if !@current_struct_def.nil?
          handle_struct(line)
        elsif !@current_enum_def.nil?
          handle_enum(line)
        elsif line[/^typedef struct ([A-z]+)\s*(\/\*.+)?$/]
          next if IGNORE_LIST.include?($1)
          @current_struct_def = ''
        elsif line[/^typedef enum( [A-z]+)?\s*(\/\*.+)?$/]
          next if IGNORE_LIST.include?($1)
          @empty_enum = true
          @current_enum_def = ''
        elsif line[/^typedef( struct)? ([A-z0-9\s_]+) \*?([A-z]+);/]
          next if IGNORE_LIST.include?($2) || IGNORE_LIST.include?($3)
          go_type = map_to_go_type($2) || $2
          @typedefs += format("type %s %s\n\n", $3, go_type)
        end
      end
    end
  end

  def handle_struct(line)
    if line[/^\s+(struct |const )?([A-z0-9]+)\s+(\*){0,2}([A-z_]+);\s*(\/\*.+)?/]
      name = $4
      c_type = $2 + $3.to_s
      comment = $5

      go_name = classify(name)
      go_type = map_to_go_type(c_type)

      if go_type
        @current_struct_def += format("\t%s\t%s\t`json:\"%s\"`\t%s\n", go_name, go_type, name, comment)
      end

      @open_comment = line.include?("/*") && !line.include?("*/")
    elsif line[/^\}\s+([A-z]+);/]
      @struct_defs[@target_group][$1] = @current_struct_def
      @current_struct_def = nil
    elsif line.strip.start_with?('/*') || @open_comment
      @current_struct_def += line
      @open_comment = !line.include?("*/")
    elsif !@current_struct_def.empty?
      @current_struct_def += "\n"
    end
  end

  def map_to_go_type(c_type)
    if c_type == 'NodeTag'
      # Ignore
    elsif ['ParamFetchHook', 'ParserSetupHook', 'FdwRoutine*'].include?(c_type)
      # Ignore (function pointers)
    elsif ['HTAB*', 'MemoryContext'].include?(c_type)
      # Ignore (can't map this properly)
    elsif c_type == 'List*'
      '[]Node'
    elsif @nodetypes.include?(c_type[0..-2])
      '*' + c_type[0..-2]
    elsif c_type == 'Node*'
      'Node'
    elsif ['bool', 'int', 'int16', 'int32', 'uint', 'uint16', 'uint32'].include?(c_type)
      c_type
    elsif ['char*', 'NameData'].include?(c_type)
      '*string'
    elsif c_type == 'char'
      'byte'
    elsif ['AttrNumber', 'Oid', 'Index'].include?(c_type)
      c_type
    elsif ['signed int', 'long'].include?(c_type)
      'int64'
    elsif c_type == 'unsigned int'
      'uint64'
    elsif c_type == 'bits32'
      'uint32'
    elsif c_type == 'float'
      'float32'
    elsif c_type == 'double'
      'float64'
    elsif ['uintptr_t', 'Size'].include?(c_type)
      'uintptr'
    elsif ['Bitmapset*', 'Bitmapset', 'Relids'].include?(c_type)
      '[]uint32'
    elsif ['Relids*', 'int32*'].include?(c_type)
      '[]' + c_type[0..-2]
    elsif c_type == 'void*'
      'interface{}'
    elsif c_type[-1] == '*'
      '*' + c_type[0..-2]
    else # Enum
      c_type
    end
  end

  def handle_enum(line)
    if line[/^\s+([A-z0-9_]+),?\s*(\/\*.+)?/]
      name = $1
      comment = $2

      if @empty_enum
        @current_enum_def += format("\t%s\t= iota\t%s\n", name, comment)
        @empty_enum = false
      else
        @current_enum_def += format("\t%s\t%s\n", name, comment)
      end

      @open_comment = line.include?("/*") && !line.include?("*/")
    elsif line[/^\}\s+([A-z]+);/]
      @enum_defs[@target_group][$1] = @current_enum_def
      @current_enum_def = nil
    elsif line.strip.start_with?('/*') || @open_comment
      @current_enum_def += line
      @open_comment = !line.include?("*/")
    elsif !@current_enum_def.empty?
      @current_enum_def += "\n"
    end
  end

  OUTNODE_NAME_OVERRIDES = {
    'VacuumStmt' => 'VACUUM',
    'InsertStmt' => 'INSERT INTO',
    'DeleteStmt' => 'DELETE FROM',
    'UpdateStmt' => 'UPDATE',
    'SelectStmt' => 'SELECT',
    'AlterTableStmt' => 'ALTER TABLE',
    'AlterTableCmd' => 'ALTER TABLE CMD',
    'CopyStmt' => 'COPY',
    'DropStmt' => 'DROP',
    'TruncateStmt' => 'TRUNCATE',
    'TransactionStmt' => 'TRANSACTION',
    'ExplainStmt' => 'EXPLAIN',
    'CreateTableAsStmt' => 'CREATE TABLE AS',
    'VariableSetStmt' => 'SET',
    'VariableShowStmt' => 'SHOW',
    'LockStmt' => 'LOCK',
    'CheckPointStmt' => 'CHECKPOINT',
    'CreateSchemaStmt' => 'CREATE SCHEMA',
    'DeclareCursorStmt' => 'DECLARECURSOR',
    'CollateExpr' => 'COLLATE',
    'CaseExpr' => 'CASE',
    'CaseWhen' => 'WHEN',
    'ArrayExpr' => 'ARRAY',
    'RowExpr' => 'ROW',
    'RowCompareExpr' => 'ROWCOMPARE',
    'CoalesceExpr' => 'COALESCE',
    'MinMaxExpr' => 'MINMAX',
    'A_Expr' => 'AEXPR',
  }

  def generate!
    generate_nodetypes!
    generate_defs!

    node_unmarshal_cases = ''

    @struct_defs.each do |group, defs|
      defs.each do |type, struct_def|
        next if IGNORE_LIST.include?(type)

        out = format("type %s struct {\n", type)
        out += struct_def
        out += "}\n"

        write_nodes_file(type, out)

        json_key = OUTNODE_NAME_OVERRIDES[type] || type.upcase

        out = %{
import "encoding/json"

func (node #{type}) MarshalJSON() ([]byte, error) {
  type #{type}MarshalAlias #{type}
  return json.Marshal(map[string]interface{}{
    "#{json_key}": (*#{type}MarshalAlias)(&node),
  })
}

func (node *#{type}) UnmarshalJSON(input []byte) (err error) {
  err = UnmarshalNodeFieldJSON(input, node)
  return
}

func (node #{type}) Deparse() string {
  panic("Not Implemented")
}
        }

        write_nodes_file(type + '_methods', out, false)

        node_unmarshal_cases << %{
\t\t\tcase "#{json_key}":
\t\t\t\tvar outNode #{type}
\t\t\t\tjson.Unmarshal(jsonText, &outNode)
\t\t\t\tnode = outNode}
      end
    end

    out_node_unmarshal = %(
import (
\t"encoding/json"
\t"strings"
)

func UnmarshalNodeJSON(input json.RawMessage) (node Node, err error) {
\t// Simple heuristic to catch value nodes
\tif (!strings.HasPrefix(string(input), "{")) {
\t\tvar value Value
\t\terr = json.Unmarshal(input, &value)
\t\tnode = value
\t\treturn
\t}

\tvar nodeMap map[string]json.RawMessage

\terr = json.Unmarshal(input, &nodeMap)
\tif err != nil {
\t\treturn
\t}

\tfor nodeType, jsonText := range nodeMap {
\t\tswitch nodeType {
#{node_unmarshal_cases}
\t\t}
\t}

\treturn
}
    )

    write_nodes_file('node_unmarshal', out_node_unmarshal, true)

    @enum_defs.each do |group, defs|
      defs.each do |type, enum_def|
        next if IGNORE_LIST.include?(type)

        out = format("type %s uint\n\n", type)
        out += "const (\n"
        out += enum_def
        out += ")\n"

        write_nodes_file(type, out)
      end
    end

    write_nodes_file('typedefs', @typedefs)
  end

  def write_nodes_file(name, content, overwrite = true)
    path = format('./nodes/%s.go', underscore(name))
    return if !overwrite && File.exist?(path)

    content = "// Auto-generated - DO NOT EDIT\n\npackage pg_query\n\n" + content
    File.write(path, content)
    system format("go fmt %s", path)
  end
end

Generator.new('./tmp/libpg_query-master/postgres').generate!
