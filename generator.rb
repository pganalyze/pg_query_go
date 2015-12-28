# rubocop:disable Style/PerlBackrefs, Metrics/AbcSize, Metrics/LineLength, Metrics/MethodLength, Style/WordArray, Metrics/ClassLength, Style/Documentation, Metrics/CyclomaticComplexity, Metrics/PerceivedComplexity, Style/TrailingComma, Style/RegexpLiteral

require 'bundler'

def underscore(str)
  str
    .gsub(/([A-Z]+)([A-Z][a-z])/, '\1_\2')
    .gsub(/([a-z\d])([A-Z])/, '\1_\2')
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
  ]

  def generate_defs!
    @target_struct = nil
    @open_comment = false

    @all_known_enums = []
    @enum_defs = {}
    @struct_defs = {}
    @typedefs = []

    ['nodes/parsenodes', 'nodes/primnodes',
     'nodes/nodes', 'nodes/params', 'access/attnum', 'c', 'postgres', 'postgres_ext',
     'storage/block', 'access/sdir'].each do |group|
      @target_group = group
      @struct_defs[@target_group] = {}
      @enum_defs[@target_group] = {}
      @comment_text = nil

      lines = File.read(File.join(@pgdir, format('/src/include/%s.h', @target_group)))
      lines.each_line do |line|
        if !@current_struct_def.nil?
          handle_struct(line)
        elsif !@current_enum_def.nil?
          handle_enum(line)
        elsif line[/^typedef struct ([A-z]+)\s*(\/\*.+)?$/]
          next if IGNORE_LIST.include?($1)
          @current_struct_def = { fields: [], comment: @open_comment_text }
          @open_comment_text = nil
        elsif line[/^typedef enum( [A-z]+)?\s*(\/\*.+)?$/]
          next if IGNORE_LIST.include?($1)
          @current_enum_def = { values: [], comment: @open_comment_text }
          @open_comment_text = nil
        elsif line[/^typedef( struct)? ([A-z0-9\s_]+) \*?([A-z]+);/]
          next if IGNORE_LIST.include?($2) || IGNORE_LIST.include?($3)
          @typedefs << { new_type_name: $3, source_type: $2, comment: @open_comment_text }
          @open_comment_text = nil
        elsif line.strip.start_with?('/*')
          @open_comment_text = line
          @open_comment = !line.include?('*/')
        elsif @open_comment
          @open_comment_text += "\n" unless @open_comment_text.end_with?("\n")
          @open_comment_text += line
          @open_comment = !line.include?('*/')
        end
      end
    end
  end

  def handle_struct(line)
    if line[/^\s+(struct |const )?([A-z0-9]+)\s+(\*){0,2}([A-z_]+);\s*(\/\*.+)?/]
      name = $4
      c_type = $2 + $3.to_s
      comment = $5

      @current_struct_def[:fields] << { name: name, c_type: c_type, comment: comment }

      @open_comment = line.include?('/*') && !line.include?('*/')
    elsif line[/^\}\s+([A-z]+);/]
      @struct_defs[@target_group][$1] = @current_struct_def
      @current_struct_def = nil
    elsif line.strip.start_with?('/*')
      @current_struct_def[:fields] << { comment: line }
      @open_comment = !line.include?('*/')
    elsif @open_comment
      @current_struct_def[:fields].last[:comment] += "\n" unless @current_struct_def[:fields].last[:comment].end_with?("\n")
      @current_struct_def[:fields].last[:comment] += line
      @open_comment = !line.include?('*/')
    elsif !@current_struct_def[:fields].empty?
      @current_struct_def[:fields] << { comment: '' }
    end
  end

  def handle_enum(line)
    if line[/^\s+([A-z0-9_]+),?\s*([A-z0-9_]+)?(\/\*.+)?/]
      name = $1
      other_name = $2
      comment = $3

      @current_enum_def[:values] << { name: name, comment: comment }
      @current_enum_def[:values] << { name: other_name } if other_name

      @open_comment = line.include?('/*') && !line.include?('*/')
    elsif line[/^\}\s+([A-z]+);/]
      @all_known_enums << $1
      @enum_defs[@target_group][$1] = @current_enum_def
      @current_enum_def = nil
    elsif line.strip.start_with?('/*')
      @current_enum_def[:values] << { comment: line }
      @open_comment = !line.include?('*/')
    elsif @open_comment
      @current_enum_def[:values].last[:comment] += "\n" unless @current_enum_def[:values].last[:comment].end_with?("\n")
      @current_enum_def[:values].last[:comment] += line
      @open_comment = !line.include?('*/')
    elsif !@current_enum_def.empty?
      @current_enum_def[:values] << { comment: '' }
    end
  end

  GO_TYPE_OVERRIDES = {
    ['SelectStmt', 'valuesLists'] => '[][]Node',
    ['Float', 'str'] => 'string',
    ['String', 'str'] => 'string',
    ['BitString', 'str'] => 'string',
  }

  def map_to_go_type(c_type)
    return if c_type == 'NodeTag' # Ignore
    return if ['ParamFetchHook', 'ParserSetupHook', 'FdwRoutine*'].include?(c_type) # Ignore (function pointers)
    return if ['HTAB*', 'MemoryContext'].include?(c_type) # Ignore (can't map this properly)

    if c_type == 'List*'
      'List'
    elsif ['Node*', 'Value', 'Expr', 'Expr*'].include?(c_type)
      'Node'
    elsif @nodetypes.include?(c_type[0..-2])
      '*' + c_type[0..-2]
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
    elsif c_type == '[]Node'
      c_type # Passthrough from an override
    else # Enum
      c_type
    end
  end

  # Top-of-struct comment special cases - we might want to merge these into the same output files at some point
  COMMENT_ENUM_TO_STRUCT = {
    'nodes/parsenodes' => {
      'SelectStmt' => 'SetOperation',
      'CreateRoleStmt' => 'RoleStmtType',
      'AlterRoleStmt' => 'RoleStmtType',
      'AlterRoleSetStmt' => 'RoleStmtType',
      'DropRoleStmt' => 'RoleStmtType',
      'A_Expr' => 'A_Expr_Kind',
      'DefElem' => 'DefElemAction',
      'DiscardStmt' => 'DiscardMode',
      'FetchStmt' => 'FetchDirection',
      'GrantStmt' => 'GrantTargetType',
      'PrivGrantee' => 'GrantTargetType',
      'LockingClause' => 'LockClauseStrength',
      'RangeTblEntry' => 'RTEKind',
      'TransactionStmt' => 'TransactionStmtKind',
      'VacuumStmt' => 'VacuumOption',
      'ViewStmt' => 'ViewCheckOption',
    },
    'nodes/primnodes' => {
      'MinMaxExpr' => 'MinMaxOp',
      'Param' => 'ParamKind',
      'RowCompareExpr' => 'RowCompareType',
      'SubLink' => 'SubLinkType',
      'BooleanTest' => 'BoolTestType',
      'NullTest' => 'NullTestType',
    }
  }
  COMMENT_STRUCT_TO_STRUCT = {
    'nodes/parsenodes' => {
      'AlterDatabaseSetStmt' => 'AlterDatabaseStmt',
      # 'AlterExtensionStmt' => 'CreateExtensionStmt', # FIXME: This overrides an existing sub-comment
      'AlterExtensionContentsStmt' => 'CreateExtensionStmt',
      'AlterFdwStmt' => 'CreateFdwStmt',
      'AlterForeignServerStmt' => 'CreateForeignServerStmt',
      'AlterFunctionStmt' => 'CreateFunctionStmt',
      'AlterSeqStmt' => 'CreateSeqStmt',
      'AlterTableCmd' => 'AlterTableStmt',
      'ReplicaIdentityStmt' => 'AlterTableStmt',
      'AlterUserMappingStmt' => 'CreateUserMappingStmt',
      'DropUserMappingStmt' => 'CreateUserMappingStmt',
      'CreateOpClassItem' => 'CreateOpClassStmt',
      'DropTableSpaceStmt' => 'CreateTableSpaceStmt',
      'FunctionParameter' => 'CreateFunctionStmt',
      'InlineCodeBlock' => 'DoStmt',
    },
    'nodes/params' => {
      'ParamListInfoData' => 'ParamExternData',
    },
  }
  def transform_toplevel_comments!
    COMMENT_ENUM_TO_STRUCT.each do |file, mapping|
      mapping.each do |target, source|
        @struct_defs[file][target][:comment] = @enum_defs[file][source][:comment]
      end
    end

    COMMENT_STRUCT_TO_STRUCT.each do |file, mapping|
      mapping.each do |target, source|
        @struct_defs[file][target][:comment] = @struct_defs[file][source][:comment]
      end
    end
  end

  LIST_FINGERPRINT = '''
  if parentFieldName == "FromClause" || parentFieldName == "TargetList" || parentFieldName == "Cols" || parentFieldName == "Rexpr" {
		var itemsFingerprints FingerprintSubContextSlice

		for _, subNode := range node.Items {
      if subNode != nil {
			  subCtx := FingerprintSubContext{}
			  subNode.Fingerprint(&subCtx, parentFieldName)
			  itemsFingerprints.AddIfUnique(subCtx)
      }
		}

		sort.Sort(itemsFingerprints)

		for _, fingerprint := range itemsFingerprints {
			for _, part := range fingerprint.parts {
				ctx.WriteString(part)
			}
		}
	} else {
    for _, subNode := range node.Items {
      if subNode != nil {
        subNode.Fingerprint(ctx, parentFieldName)
      }
    }
  }
  '''

  # Fingerprinting additional code to be inserted
  FINGERPRINT_OVERRIDE_NODES = {
    'A_Const' => :skip,
    'Alias' => :skip,
    'ParamRef' => :skip,
    'SetToDefault' => :skip,
    'List' => LIST_FINGERPRINT,
  }
  FINGERPRINT_OVERRIDE_FIELDS = {
    [nil, 'location'] => :skip,
    ['ResTarget', 'name'] => "if node.Name != nil && parentFieldName != \"TargetList\" {\nctx.WriteString(*node.Name)\n}\n",
    ['PrepareStmt', 'name'] => :skip,
    ['ExecuteStmt', 'name'] => :skip,
    ['DeallocateStmt', 'name'] => :skip,
  }
  GO_INT_TYPES = ['int', 'int16', 'int32', 'int64', 'uint16', 'uint32', 'uint64', 'Oid', 'Index', 'AclMode', 'AttrNumber']
  GO_INT_ARRAY_TYPES = ['[]uint32']
  GO_FLOAT_TYPES = ['Cost']

  def generate!
    generate_nodetypes!
    generate_defs!
    transform_toplevel_comments!

    node_unmarshal_cases = []

    @struct_defs['nodes/value'] = {}
    @struct_defs['nodes/value']['Integer'] = { fields: [{ name: 'ival', c_type: 'long' }] }
    @struct_defs['nodes/value']['Float'] = { fields: [{ name: 'str', c_type: 'char*' }] }
    @struct_defs['nodes/value']['String'] = { fields: [{ name: 'str', c_type: 'char*' }] }
    @struct_defs['nodes/value']['BitString'] = { fields: [{ name: 'str', c_type: 'char*' }] }
    @struct_defs['nodes/pg_list'] = { 'List' => { fields: [{ name: 'items', c_type: '[]Node' }] } }
    @struct_defs['nodes/value']['Null'] = { fields: [] }

    @struct_defs.each do |source_filename, defs|
      defs.each do |type, struct_def|
        next if IGNORE_LIST.include?(type)

        struct_def_go = ''
        unmarshal_def = ''
        struct_def[:fields].each_with_index do |field, index|
          if !field[:name] && field[:comment]
            struct_def_go += "\n" if index != 0
            struct_def_go += field[:comment]
            next
          end

          go_name = classify(field[:name])
          go_type = GO_TYPE_OVERRIDES[[type, field[:name]]] || map_to_go_type(field[:c_type])
          next unless go_type

          struct_def_go += format("%s\t%s\t`json:\"%s\"`\t%s\n", go_name, go_type, field[:name], field[:comment])
          unmarshal_def += format("if fields[\"%s\"] != nil {\n", field[:name])

          if go_type == '[][]Node'
            unmarshal_def += format("node.%s, err = UnmarshalNodeArrayArrayJSON(fields[\"%s\"])\n", go_name, field[:name])
            unmarshal_def += "if err != nil {\nreturn\n}\n"
          elsif go_type == '[]Node'
            unmarshal_def += format("node.%s, err = UnmarshalNodeArrayJSON(fields[\"%s\"])\n", go_name, field[:name])
            unmarshal_def += "if err != nil {\nreturn\n}\n"
          elsif go_type == 'Node'
            unmarshal_def += format("node.%s, err = UnmarshalNodeJSON(fields[\"%s\"])\n", go_name, field[:name])
            unmarshal_def += "if err != nil {\nreturn\n}\n"
          elsif go_type == 'List'
            unmarshal_def += format("node.%s.Items, err = UnmarshalNodeArrayJSON(fields[\"%s\"])\n", go_name, field[:name])
            unmarshal_def += "if err != nil {\nreturn\n}\n"
          elsif go_type[0].start_with?('*') && @nodetypes.include?(go_type[1..-1])
            unmarshal_def += "var nodePtr *Node\n"
            unmarshal_def += format("nodePtr, err = UnmarshalNodePtrJSON(fields[\"%s\"])\n", field[:name])
            unmarshal_def += "if err != nil {\nreturn\n}\n"
            unmarshal_def += "if nodePtr != nil && *nodePtr != nil {\n"
            unmarshal_def += format("val := (*nodePtr).(%s)\n", go_type[1..-1])
            unmarshal_def += format("node.%s = &val\n", go_name)
            unmarshal_def += "}\n"
          elsif go_type == 'byte'
            unmarshal_def += "var strVal string\n"
            unmarshal_def += format("err = json.Unmarshal(fields[\"%s\"], &strVal)\n", field[:name])
            unmarshal_def += format("node.%s = strVal[0]\n", go_name)
            unmarshal_def += "if err != nil {\nreturn\n}\n"
          else
            unmarshal_def += format("err = json.Unmarshal(fields[\"%s\"], &node.%s)\n", field[:name], go_name)
            unmarshal_def += "if err != nil {\nreturn\n}\n"
          end
          unmarshal_def += "}\n\n"
        end

        fp_override = FINGERPRINT_OVERRIDE_NODES[type]
        if fp_override
          fp_override = '// Intentionally ignoring all fields for fingerprinting' if fp_override == :skip
          fingerprint_def = fp_override
        else
          fingerprint_def = format("ctx.WriteString(\"%s\")\n", type)
          struct_def[:fields].reject { |f| f[:name].nil? }.sort_by { |f| f[:name] }.each do |field|
            go_name = classify(field[:name])
            go_type = GO_TYPE_OVERRIDES[[type, field[:name]]] || map_to_go_type(field[:c_type])
            next unless go_type

            fp_override = FINGERPRINT_OVERRIDE_FIELDS[[type, field[:name]]] || FINGERPRINT_OVERRIDE_FIELDS[[nil, field[:name]]]
            if fp_override
              if fp_override == :skip
                fp_override = format('// Intentionally ignoring node.%s for fingerprinting', go_name)
              end
              fingerprint_def += fp_override + "\n\n"
              next
            end

            case go_type
            when '[][]Node'
              fingerprint_def += format("if len(node.%s) > 0 {\n", go_name)
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("for _, nodeList := range node.%s {\n", go_name)
              fingerprint_def += "for _, subNode := range nodeList {\n"
              fingerprint_def += format("subNode.Fingerprint(ctx, \"%s\")\n", go_name)
              fingerprint_def += "}\n"
              fingerprint_def += "}\n"
              fingerprint_def += "}\n\n"
            when '[]Node'
              fingerprint_def += format("if len(node.%s) > 0 {\n", go_name)
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("for _, subNode := range node.%s {\n", go_name)
              fingerprint_def += format("subNode.Fingerprint(ctx, \"%s\")\n", go_name)
              fingerprint_def += "}\n"
              fingerprint_def += "}\n\n"
            when 'Node'
              fingerprint_def += format("\nif node.%s != nil {", go_name)
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("node.%s.Fingerprint(ctx, \"%s\")\n", go_name, go_name)
              fingerprint_def += "}\n\n"
            when 'List'
              fingerprint_def += format("if len(node.%s.Items) > 0 {\n", go_name)
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("node.%s.Fingerprint(ctx, \"%s\")\n", go_name, go_name)
              fingerprint_def += "}\n\n"
            when 'CreateStmt'
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("node.%s.Fingerprint(ctx, \"%s\")\n\n", go_name, go_name)
            when 'byte'
              fingerprint_def += format("\nif node.%s != 0 {", go_name)
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("ctx.WriteString(string(node.%s))\n\n", go_name)
              fingerprint_def += "}\n\n"
            when 'string'
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("ctx.WriteString(node.%s)\n\n", go_name)
            when '*string'
              fingerprint_def += format("\nif node.%s != nil {", go_name)
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("ctx.WriteString(*node.%s)\n", go_name)
              fingerprint_def += "}\n\n"
            when 'bool'
              fingerprint_def += format("\nif node.%s {", go_name)
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("ctx.WriteString(strconv.FormatBool(node.%s))\n", go_name)
              fingerprint_def += "}\n\n"
            when 'Datum', 'interface{}'
              # Ignore
            when *GO_INT_TYPES
              fingerprint_def += format("\nif node.%s != 0 {", go_name)
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("ctx.WriteString(strconv.Itoa(int(node.%s)))\n", go_name)
              fingerprint_def += "}\n\n"
            when *GO_INT_ARRAY_TYPES
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("for _, val := range node.%s {\n", go_name)
              fingerprint_def += "ctx.WriteString(strconv.Itoa(int(val)))\n"
              fingerprint_def += "}\n\n"
            when *GO_FLOAT_TYPES
              fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
              fingerprint_def += format("ctx.WriteString(strconv.FormatFloat(float64(node.%s), 'E', -1, 64))\n\n", go_name)
            else
              if go_type[0].start_with?('*') && @nodetypes.include?(go_type[1..-1])
                fingerprint_def += format("\nif node.%s != nil {", go_name)
                fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
                fingerprint_def += format("node.%s.Fingerprint(ctx, \"%s\")\n", go_name, go_name)
                fingerprint_def += "}\n\n"
              elsif @all_known_enums.include?(go_type)
                fingerprint_def += format("\nif int(node.%s) != 0 {", go_name)
                fingerprint_def += format("ctx.WriteString(\"%s\")\n", field[:name])
                fingerprint_def += format("ctx.WriteString(strconv.Itoa(int(node.%s)))\n", go_name)
                fingerprint_def += "}\n\n"
              else
                # This shouldn't happen - if it does the above is missing something :-)
                puts type
                puts go_name
                fail go_type
              end
            end
          end
        end

        write_nodes_file type, %(
import "encoding/json"

#{struct_def[:comment] && struct_def[:comment].strip}
type #{type} struct {
#{struct_def_go.strip}
}

func (node #{type}) MarshalJSON() ([]byte, error) {
  type #{type}MarshalAlias #{type}
  return json.Marshal(map[string]interface{}{
    "#{type}": (*#{type}MarshalAlias)(&node),
  })
}

func (node *#{type}) UnmarshalJSON(input []byte) (err error) {
  var fields map[string]json.RawMessage

  err = json.Unmarshal(input, &fields)
  if err != nil {
    return
  }

  #{unmarshal_def}

  return
}
        ), true, "postgres/src/include/#{source_filename}.h"

        write_nodes_file(type + '_deparse', %(
func (node #{type}) Deparse() string {
  panic("Not Implemented")
}
        ), true)

        write_nodes_file(type + '_fingerprint', %(
import "strconv"

func (node #{type}) Fingerprint(ctx FingerprintContext, parentFieldName string) {
  #{fingerprint_def.strip}
}
        ), true)

        node_unmarshal_cases << type
      end
    end

    node_unmarshal_go = ''
    node_unmarshal_cases.each do |type|
      node_unmarshal_go += %(
      case "#{type}":
      var outNode #{type}
      err = json.Unmarshal(jsonText, &outNode)
      if err != nil {
        return
      }
      node = outNode)
    end

    write_nodes_file 'node_unmarshal', %(
import (
  "encoding/json"
  "strings"
)

func UnmarshalNodeJSON(input json.RawMessage) (node Node, err error) {
  if input == nil || string(input) == "null" {
    return
  }

  if strings.HasPrefix(string(input), "[") {
		var list List
		list.Items, err = UnmarshalNodeArrayJSON(input)
		if err != nil {
			return
		}
		node = list
		return
	}

  var nodeMap map[string]json.RawMessage

  err = json.Unmarshal(input, &nodeMap)
  if err != nil {
    return
  }

  for nodeType, jsonText := range nodeMap {
    switch nodeType {
#{node_unmarshal_go}
    }
  }

  return
}
    )

    @enum_defs.each do |source_filename, defs|
      defs.each do |type, enum_def|
        next if IGNORE_LIST.include?(type)

        go_enum_def = ''
        output_first_type_field = false
        enum_def[:values].each_with_index do |field, index|
          if !field[:name] && field[:comment]
            go_enum_def += "\n" if index != 0
            go_enum_def += field[:comment]
            next
          end

          if !output_first_type_field
            go_enum_def += format("%s %s = iota %s\n", field[:name], type, field[:comment])
            output_first_type_field = true
          else
            go_enum_def += format("%s\t%s\n", field[:name], field[:comment])
          end
        end

        write_nodes_file type, %(
          #{enum_def[:comment] && enum_def[:comment].strip}
          type #{type} uint

          const (
            #{go_enum_def.strip}
          )
        ), true, "postgres/src/include/#{source_filename}.h"
      end
    end

    typedefs_go = ''
    @typedefs.each do |typedef|
      go_type = map_to_go_type(typedef[:source_type]) || typedef[:source_type]
      typedefs_go += format("type %s %s\n\n", typedef[:new_type_name], go_type)
    end

    write_nodes_file('typedefs', typedefs_go)
  end

  def write_nodes_file(name, content, overwrite = true, source_file = nil)
    name += '_expr' if name.end_with?('Test')
    path = format('./nodes/%s.go', underscore(name))
    return if !overwrite && File.exist?(path)

    if source_file
      header = "// Auto-generated from #{source_file} - DO NOT EDIT\n"
    else
      header = "// Auto-generated - DO NOT EDIT\n"
    end
    content = header + "\npackage pg_query\n\n" + content
    File.write(path, content)
    system format('go fmt %s', path)
    system format('goimports -w %s', path)
  end
end

Generator.new('./tmp/libpg_query-master/postgres').generate!
