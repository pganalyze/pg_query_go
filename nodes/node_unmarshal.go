// Auto-generated - DO NOT EDIT

package pg_query

import (
	"encoding/json"
	"strings"
)

func UnmarshalNodeJSON(input json.RawMessage) (node Node, err error) {
	// Simple heuristic to catch value nodes
	if !strings.HasPrefix(string(input), "{") {
		var value Value
		err = json.Unmarshal(input, &value)
		node = value
		return
	}

	var nodeMap map[string]json.RawMessage

	err = json.Unmarshal(input, &nodeMap)
	if err != nil {
		return
	}

	for nodeType, jsonText := range nodeMap {
		switch nodeType {

		case "QUERY":
			var outNode Query
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "TYPENAME":
			var outNode TypeName
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COLUMNREF":
			var outNode ColumnRef
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PARAMREF":
			var outNode ParamRef
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "AEXPR":
			var outNode A_Expr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "A_CONST":
			var outNode A_Const
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "TYPECAST":
			var outNode TypeCast
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COLLATECLAUSE":
			var outNode CollateClause
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FUNCCALL":
			var outNode FuncCall
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "A_STAR":
			var outNode A_Star
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "A_INDICES":
			var outNode A_Indices
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "A_INDIRECTION":
			var outNode A_Indirection
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "A_ARRAYEXPR":
			var outNode A_ArrayExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RESTARGET":
			var outNode ResTarget
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SORTBY":
			var outNode SortBy
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "WINDOWDEF":
			var outNode WindowDef
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RANGESUBSELECT":
			var outNode RangeSubselect
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RANGEFUNCTION":
			var outNode RangeFunction
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COLUMNDEF":
			var outNode ColumnDef
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "TABLELIKECLAUSE":
			var outNode TableLikeClause
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "INDEXELEM":
			var outNode IndexElem
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DEFELEM":
			var outNode DefElem
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "LOCKINGCLAUSE":
			var outNode LockingClause
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "XMLSERIALIZE":
			var outNode XmlSerialize
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RANGETBLENTRY":
			var outNode RangeTblEntry
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RANGETBLFUNCTION":
			var outNode RangeTblFunction
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "WITHCHECKOPTION":
			var outNode WithCheckOption
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SORTGROUPCLAUSE":
			var outNode SortGroupClause
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "WINDOWCLAUSE":
			var outNode WindowClause
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ROWMARKCLAUSE":
			var outNode RowMarkClause
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "WITHCLAUSE":
			var outNode WithClause
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COMMONTABLEEXPR":
			var outNode CommonTableExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "INSERT INTO":
			var outNode InsertStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DELETE FROM":
			var outNode DeleteStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "UPDATE":
			var outNode UpdateStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SELECT":
			var outNode SelectStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SETOPERATIONSTMT":
			var outNode SetOperationStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATE SCHEMA":
			var outNode CreateSchemaStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTER TABLE":
			var outNode AlterTableStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "REPLICAIDENTITYSTMT":
			var outNode ReplicaIdentityStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTER TABLE CMD":
			var outNode AlterTableCmd
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERDOMAINSTMT":
			var outNode AlterDomainStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "GRANTSTMT":
			var outNode GrantStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PRIVGRANTEE":
			var outNode PrivGrantee
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FUNCWITHARGS":
			var outNode FuncWithArgs
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ACCESSPRIV":
			var outNode AccessPriv
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "GRANTROLESTMT":
			var outNode GrantRoleStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERDEFAULTPRIVILEGESSTMT":
			var outNode AlterDefaultPrivilegesStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COPY":
			var outNode CopyStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SET":
			var outNode VariableSetStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SHOW":
			var outNode VariableShowStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATESTMT":
			var outNode CreateStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CONSTRAINT":
			var outNode Constraint
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATETABLESPACESTMT":
			var outNode CreateTableSpaceStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DROPTABLESPACESTMT":
			var outNode DropTableSpaceStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERTABLESPACEOPTIONSSTMT":
			var outNode AlterTableSpaceOptionsStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERTABLEMOVEALLSTMT":
			var outNode AlterTableMoveAllStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEEXTENSIONSTMT":
			var outNode CreateExtensionStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTEREXTENSIONSTMT":
			var outNode AlterExtensionStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTEREXTENSIONCONTENTSSTMT":
			var outNode AlterExtensionContentsStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEFDWSTMT":
			var outNode CreateFdwStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERFDWSTMT":
			var outNode AlterFdwStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEFOREIGNSERVERSTMT":
			var outNode CreateForeignServerStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERFOREIGNSERVERSTMT":
			var outNode AlterForeignServerStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEFOREIGNTABLESTMT":
			var outNode CreateForeignTableStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEUSERMAPPINGSTMT":
			var outNode CreateUserMappingStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERUSERMAPPINGSTMT":
			var outNode AlterUserMappingStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DROPUSERMAPPINGSTMT":
			var outNode DropUserMappingStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATETRIGSTMT":
			var outNode CreateTrigStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEEVENTTRIGSTMT":
			var outNode CreateEventTrigStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTEREVENTTRIGSTMT":
			var outNode AlterEventTrigStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEPLANGSTMT":
			var outNode CreatePLangStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEROLESTMT":
			var outNode CreateRoleStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERROLESTMT":
			var outNode AlterRoleStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERROLESETSTMT":
			var outNode AlterRoleSetStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DROPROLESTMT":
			var outNode DropRoleStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATESEQSTMT":
			var outNode CreateSeqStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERSEQSTMT":
			var outNode AlterSeqStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DEFINESTMT":
			var outNode DefineStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEDOMAINSTMT":
			var outNode CreateDomainStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEOPCLASSSTMT":
			var outNode CreateOpClassStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEOPCLASSITEM":
			var outNode CreateOpClassItem
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEOPFAMILYSTMT":
			var outNode CreateOpFamilyStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTEROPFAMILYSTMT":
			var outNode AlterOpFamilyStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DROP":
			var outNode DropStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "TRUNCATE":
			var outNode TruncateStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COMMENTSTMT":
			var outNode CommentStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SECLABELSTMT":
			var outNode SecLabelStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DECLARECURSOR":
			var outNode DeclareCursorStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CLOSEPORTALSTMT":
			var outNode ClosePortalStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FETCHSTMT":
			var outNode FetchStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "INDEXSTMT":
			var outNode IndexStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEFUNCTIONSTMT":
			var outNode CreateFunctionStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FUNCTIONPARAMETER":
			var outNode FunctionParameter
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERFUNCTIONSTMT":
			var outNode AlterFunctionStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DOSTMT":
			var outNode DoStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "INLINECODEBLOCK":
			var outNode InlineCodeBlock
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RENAMESTMT":
			var outNode RenameStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTEROBJECTSCHEMASTMT":
			var outNode AlterObjectSchemaStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTEROWNERSTMT":
			var outNode AlterOwnerStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RULESTMT":
			var outNode RuleStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "NOTIFYSTMT":
			var outNode NotifyStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "LISTENSTMT":
			var outNode ListenStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "UNLISTENSTMT":
			var outNode UnlistenStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "TRANSACTION":
			var outNode TransactionStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COMPOSITETYPESTMT":
			var outNode CompositeTypeStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEENUMSTMT":
			var outNode CreateEnumStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATERANGESTMT":
			var outNode CreateRangeStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERENUMSTMT":
			var outNode AlterEnumStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "VIEWSTMT":
			var outNode ViewStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "LOADSTMT":
			var outNode LoadStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATEDBSTMT":
			var outNode CreatedbStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERDATABASESTMT":
			var outNode AlterDatabaseStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERDATABASESETSTMT":
			var outNode AlterDatabaseSetStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DROPDBSTMT":
			var outNode DropdbStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERSYSTEMSTMT":
			var outNode AlterSystemStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CLUSTERSTMT":
			var outNode ClusterStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "VACUUM":
			var outNode VacuumStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "EXPLAIN":
			var outNode ExplainStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATE TABLE AS":
			var outNode CreateTableAsStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "REFRESHMATVIEWSTMT":
			var outNode RefreshMatViewStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CHECKPOINT":
			var outNode CheckPointStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DISCARDSTMT":
			var outNode DiscardStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "LOCK":
			var outNode LockStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CONSTRAINTSSETSTMT":
			var outNode ConstraintsSetStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "REINDEXSTMT":
			var outNode ReindexStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATECONVERSIONSTMT":
			var outNode CreateConversionStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CREATECASTSTMT":
			var outNode CreateCastStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PREPARESTMT":
			var outNode PrepareStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "EXECUTESTMT":
			var outNode ExecuteStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DEALLOCATESTMT":
			var outNode DeallocateStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "DROPOWNEDSTMT":
			var outNode DropOwnedStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "REASSIGNOWNEDSTMT":
			var outNode ReassignOwnedStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERTSDICTIONARYSTMT":
			var outNode AlterTSDictionaryStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERTSCONFIGURATIONSTMT":
			var outNode AlterTSConfigurationStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PLANNEDSTMT":
			var outNode PlannedStmt
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PLAN":
			var outNode Plan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RESULT":
			var outNode Result
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MODIFYTABLE":
			var outNode ModifyTable
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "APPEND":
			var outNode Append
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MERGEAPPEND":
			var outNode MergeAppend
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RECURSIVEUNION":
			var outNode RecursiveUnion
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "BITMAPAND":
			var outNode BitmapAnd
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "BITMAPOR":
			var outNode BitmapOr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SCAN":
			var outNode Scan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "INDEXSCAN":
			var outNode IndexScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "INDEXONLYSCAN":
			var outNode IndexOnlyScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "BITMAPINDEXSCAN":
			var outNode BitmapIndexScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "BITMAPHEAPSCAN":
			var outNode BitmapHeapScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "TIDSCAN":
			var outNode TidScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SUBQUERYSCAN":
			var outNode SubqueryScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FUNCTIONSCAN":
			var outNode FunctionScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "VALUESSCAN":
			var outNode ValuesScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CTESCAN":
			var outNode CteScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "WORKTABLESCAN":
			var outNode WorkTableScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FOREIGNSCAN":
			var outNode ForeignScan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "JOIN":
			var outNode Join
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "NESTLOOP":
			var outNode NestLoop
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "NESTLOOPPARAM":
			var outNode NestLoopParam
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MERGEJOIN":
			var outNode MergeJoin
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "HASHJOIN":
			var outNode HashJoin
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MATERIAL":
			var outNode Material
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SORT":
			var outNode Sort
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "GROUP":
			var outNode Group
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "AGG":
			var outNode Agg
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "WINDOWAGG":
			var outNode WindowAgg
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "UNIQUE":
			var outNode Unique
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "HASH":
			var outNode Hash
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SETOP":
			var outNode SetOp
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "LOCKROWS":
			var outNode LockRows
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "LIMIT":
			var outNode Limit
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PLANROWMARK":
			var outNode PlanRowMark
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PLANINVALITEM":
			var outNode PlanInvalItem
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALIAS":
			var outNode Alias
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RANGEVAR":
			var outNode RangeVar
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "INTOCLAUSE":
			var outNode IntoClause
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "EXPR":
			var outNode Expr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "VAR":
			var outNode Var
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CONST":
			var outNode Const
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PARAM":
			var outNode Param
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "AGGREF":
			var outNode Aggref
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "WINDOWFUNC":
			var outNode WindowFunc
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ARRAYREF":
			var outNode ArrayRef
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FUNCEXPR":
			var outNode FuncExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "NAMEDARGEXPR":
			var outNode NamedArgExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "OPEXPR":
			var outNode OpExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SCALARARRAYOPEXPR":
			var outNode ScalarArrayOpExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "BOOLEXPR":
			var outNode BoolExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SUBLINK":
			var outNode SubLink
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SUBPLAN":
			var outNode SubPlan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ALTERNATIVESUBPLAN":
			var outNode AlternativeSubPlan
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FIELDSELECT":
			var outNode FieldSelect
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FIELDSTORE":
			var outNode FieldStore
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RELABELTYPE":
			var outNode RelabelType
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COERCEVIAIO":
			var outNode CoerceViaIO
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ARRAYCOERCEEXPR":
			var outNode ArrayCoerceExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CONVERTROWTYPEEXPR":
			var outNode ConvertRowtypeExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COLLATE":
			var outNode CollateExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CASE":
			var outNode CaseExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "WHEN":
			var outNode CaseWhen
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CASETESTEXPR":
			var outNode CaseTestExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ARRAY":
			var outNode ArrayExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ROW":
			var outNode RowExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "ROWCOMPARE":
			var outNode RowCompareExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COALESCE":
			var outNode CoalesceExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MINMAX":
			var outNode MinMaxExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "XMLEXPR":
			var outNode XmlExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COERCETODOMAIN":
			var outNode CoerceToDomain
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "COERCETODOMAINVALUE":
			var outNode CoerceToDomainValue
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SETTODEFAULT":
			var outNode SetToDefault
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "CURRENTOFEXPR":
			var outNode CurrentOfExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "TARGETENTRY":
			var outNode TargetEntry
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RANGETBLREF":
			var outNode RangeTblRef
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "JOINEXPR":
			var outNode JoinExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FROMEXPR":
			var outNode FromExpr
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "QUALCOST":
			var outNode QualCost
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "AGGCLAUSECOSTS":
			var outNode AggClauseCosts
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PLANNERGLOBAL":
			var outNode PlannerGlobal
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PLANNERINFO":
			var outNode PlannerInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RELOPTINFO":
			var outNode RelOptInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "INDEXOPTINFO":
			var outNode IndexOptInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "EQUIVALENCECLASS":
			var outNode EquivalenceClass
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "EQUIVALENCEMEMBER":
			var outNode EquivalenceMember
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PATHKEY":
			var outNode PathKey
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PARAMPATHINFO":
			var outNode ParamPathInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PATH":
			var outNode Path
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "INDEXPATH":
			var outNode IndexPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "BITMAPHEAPPATH":
			var outNode BitmapHeapPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "BITMAPANDPATH":
			var outNode BitmapAndPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "BITMAPORPATH":
			var outNode BitmapOrPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "TIDPATH":
			var outNode TidPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "FOREIGNPATH":
			var outNode ForeignPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "APPENDPATH":
			var outNode AppendPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MERGEAPPENDPATH":
			var outNode MergeAppendPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RESULTPATH":
			var outNode ResultPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MATERIALPATH":
			var outNode MaterialPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "UNIQUEPATH":
			var outNode UniquePath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "JOINPATH":
			var outNode JoinPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MERGEPATH":
			var outNode MergePath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "HASHPATH":
			var outNode HashPath
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "RESTRICTINFO":
			var outNode RestrictInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MERGESCANSELCACHE":
			var outNode MergeScanSelCache
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PLACEHOLDERVAR":
			var outNode PlaceHolderVar
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SPECIALJOININFO":
			var outNode SpecialJoinInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "LATERALJOININFO":
			var outNode LateralJoinInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "APPENDRELINFO":
			var outNode AppendRelInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PLACEHOLDERINFO":
			var outNode PlaceHolderInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "MINMAXAGGINFO":
			var outNode MinMaxAggInfo
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PLANNERPARAMITEM":
			var outNode PlannerParamItem
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "SEMIANTIJOINFACTORS":
			var outNode SemiAntiJoinFactors
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "JOINCOSTWORKSPACE":
			var outNode JoinCostWorkspace
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PARAMEXTERNDATA":
			var outNode ParamExternData
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PARAMLISTINFODATA":
			var outNode ParamListInfoData
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "PARAMEXECDATA":
			var outNode ParamExecData
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "VARATT_EXTERNAL":
			var outNode varatt_external
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		case "BLOCKIDDATA":
			var outNode BlockIdData
			json.Unmarshal(jsonText, &outNode)
			node = outNode
		}
	}

	return
}
