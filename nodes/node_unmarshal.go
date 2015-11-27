// Auto-generated - DO NOT EDIT

package pg_query

import (
	"encoding/json"
	"strings"
)

func UnmarshalNodeJSON(input json.RawMessage) (node Node, err error) {
	if input == nil || string(input) == "null" {
		return
	}

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
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "TYPENAME":
			var outNode TypeName
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COLUMNREF":
			var outNode ColumnRef
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PARAMREF":
			var outNode ParamRef
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "AEXPR":
			var outNode A_Expr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "TYPECAST":
			var outNode TypeCast
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COLLATECLAUSE":
			var outNode CollateClause
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FUNCCALL":
			var outNode FuncCall
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "A_STAR":
			var outNode A_Star
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "A_INDICES":
			var outNode A_Indices
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "A_INDIRECTION":
			var outNode A_Indirection
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "A_ARRAYEXPR":
			var outNode A_ArrayExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RESTARGET":
			var outNode ResTarget
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SORTBY":
			var outNode SortBy
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "WINDOWDEF":
			var outNode WindowDef
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RANGESUBSELECT":
			var outNode RangeSubselect
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RANGEFUNCTION":
			var outNode RangeFunction
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COLUMNDEF":
			var outNode ColumnDef
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "TABLELIKECLAUSE":
			var outNode TableLikeClause
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "INDEXELEM":
			var outNode IndexElem
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DEFELEM":
			var outNode DefElem
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "LOCKINGCLAUSE":
			var outNode LockingClause
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "XMLSERIALIZE":
			var outNode XmlSerialize
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RANGETBLENTRY":
			var outNode RangeTblEntry
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RANGETBLFUNCTION":
			var outNode RangeTblFunction
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "WITHCHECKOPTION":
			var outNode WithCheckOption
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SORTGROUPCLAUSE":
			var outNode SortGroupClause
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "WINDOWCLAUSE":
			var outNode WindowClause
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ROWMARKCLAUSE":
			var outNode RowMarkClause
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "WITHCLAUSE":
			var outNode WithClause
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COMMONTABLEEXPR":
			var outNode CommonTableExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "INSERT INTO":
			var outNode InsertStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DELETE FROM":
			var outNode DeleteStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "UPDATE":
			var outNode UpdateStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SELECT":
			var outNode SelectStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SETOPERATIONSTMT":
			var outNode SetOperationStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATE SCHEMA":
			var outNode CreateSchemaStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTER TABLE":
			var outNode AlterTableStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "REPLICAIDENTITYSTMT":
			var outNode ReplicaIdentityStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTER TABLE CMD":
			var outNode AlterTableCmd
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERDOMAINSTMT":
			var outNode AlterDomainStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "GRANTSTMT":
			var outNode GrantStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PRIVGRANTEE":
			var outNode PrivGrantee
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FUNCWITHARGS":
			var outNode FuncWithArgs
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ACCESSPRIV":
			var outNode AccessPriv
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "GRANTROLESTMT":
			var outNode GrantRoleStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERDEFAULTPRIVILEGESSTMT":
			var outNode AlterDefaultPrivilegesStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COPY":
			var outNode CopyStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SET":
			var outNode VariableSetStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SHOW":
			var outNode VariableShowStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATESTMT":
			var outNode CreateStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CONSTRAINT":
			var outNode Constraint
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATETABLESPACESTMT":
			var outNode CreateTableSpaceStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DROPTABLESPACESTMT":
			var outNode DropTableSpaceStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERTABLESPACEOPTIONSSTMT":
			var outNode AlterTableSpaceOptionsStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERTABLEMOVEALLSTMT":
			var outNode AlterTableMoveAllStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEEXTENSIONSTMT":
			var outNode CreateExtensionStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTEREXTENSIONSTMT":
			var outNode AlterExtensionStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTEREXTENSIONCONTENTSSTMT":
			var outNode AlterExtensionContentsStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEFDWSTMT":
			var outNode CreateFdwStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERFDWSTMT":
			var outNode AlterFdwStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEFOREIGNSERVERSTMT":
			var outNode CreateForeignServerStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERFOREIGNSERVERSTMT":
			var outNode AlterForeignServerStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEFOREIGNTABLESTMT":
			var outNode CreateForeignTableStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEUSERMAPPINGSTMT":
			var outNode CreateUserMappingStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERUSERMAPPINGSTMT":
			var outNode AlterUserMappingStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DROPUSERMAPPINGSTMT":
			var outNode DropUserMappingStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATETRIGSTMT":
			var outNode CreateTrigStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEEVENTTRIGSTMT":
			var outNode CreateEventTrigStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTEREVENTTRIGSTMT":
			var outNode AlterEventTrigStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEPLANGSTMT":
			var outNode CreatePLangStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEROLESTMT":
			var outNode CreateRoleStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERROLESTMT":
			var outNode AlterRoleStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERROLESETSTMT":
			var outNode AlterRoleSetStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DROPROLESTMT":
			var outNode DropRoleStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATESEQSTMT":
			var outNode CreateSeqStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERSEQSTMT":
			var outNode AlterSeqStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DEFINESTMT":
			var outNode DefineStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEDOMAINSTMT":
			var outNode CreateDomainStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEOPCLASSSTMT":
			var outNode CreateOpClassStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEOPCLASSITEM":
			var outNode CreateOpClassItem
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEOPFAMILYSTMT":
			var outNode CreateOpFamilyStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTEROPFAMILYSTMT":
			var outNode AlterOpFamilyStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DROP":
			var outNode DropStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "TRUNCATE":
			var outNode TruncateStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COMMENTSTMT":
			var outNode CommentStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SECLABELSTMT":
			var outNode SecLabelStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DECLARECURSOR":
			var outNode DeclareCursorStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CLOSEPORTALSTMT":
			var outNode ClosePortalStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FETCHSTMT":
			var outNode FetchStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "INDEXSTMT":
			var outNode IndexStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEFUNCTIONSTMT":
			var outNode CreateFunctionStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FUNCTIONPARAMETER":
			var outNode FunctionParameter
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERFUNCTIONSTMT":
			var outNode AlterFunctionStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DOSTMT":
			var outNode DoStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "INLINECODEBLOCK":
			var outNode InlineCodeBlock
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RENAMESTMT":
			var outNode RenameStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTEROBJECTSCHEMASTMT":
			var outNode AlterObjectSchemaStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTEROWNERSTMT":
			var outNode AlterOwnerStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RULESTMT":
			var outNode RuleStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "NOTIFYSTMT":
			var outNode NotifyStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "LISTENSTMT":
			var outNode ListenStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "UNLISTENSTMT":
			var outNode UnlistenStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "TRANSACTION":
			var outNode TransactionStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COMPOSITETYPESTMT":
			var outNode CompositeTypeStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEENUMSTMT":
			var outNode CreateEnumStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATERANGESTMT":
			var outNode CreateRangeStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERENUMSTMT":
			var outNode AlterEnumStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "VIEWSTMT":
			var outNode ViewStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "LOADSTMT":
			var outNode LoadStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATEDBSTMT":
			var outNode CreatedbStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERDATABASESTMT":
			var outNode AlterDatabaseStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERDATABASESETSTMT":
			var outNode AlterDatabaseSetStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DROPDBSTMT":
			var outNode DropdbStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERSYSTEMSTMT":
			var outNode AlterSystemStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CLUSTERSTMT":
			var outNode ClusterStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "VACUUM":
			var outNode VacuumStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "EXPLAIN":
			var outNode ExplainStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATE TABLE AS":
			var outNode CreateTableAsStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "REFRESHMATVIEWSTMT":
			var outNode RefreshMatViewStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CHECKPOINT":
			var outNode CheckPointStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DISCARDSTMT":
			var outNode DiscardStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "LOCK":
			var outNode LockStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CONSTRAINTSSETSTMT":
			var outNode ConstraintsSetStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "REINDEXSTMT":
			var outNode ReindexStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATECONVERSIONSTMT":
			var outNode CreateConversionStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CREATECASTSTMT":
			var outNode CreateCastStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PREPARESTMT":
			var outNode PrepareStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "EXECUTESTMT":
			var outNode ExecuteStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DEALLOCATESTMT":
			var outNode DeallocateStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "DROPOWNEDSTMT":
			var outNode DropOwnedStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "REASSIGNOWNEDSTMT":
			var outNode ReassignOwnedStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERTSDICTIONARYSTMT":
			var outNode AlterTSDictionaryStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERTSCONFIGURATIONSTMT":
			var outNode AlterTSConfigurationStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PLANNEDSTMT":
			var outNode PlannedStmt
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PLAN":
			var outNode Plan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RESULT":
			var outNode Result
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MODIFYTABLE":
			var outNode ModifyTable
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "APPEND":
			var outNode Append
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MERGEAPPEND":
			var outNode MergeAppend
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RECURSIVEUNION":
			var outNode RecursiveUnion
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "BITMAPAND":
			var outNode BitmapAnd
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "BITMAPOR":
			var outNode BitmapOr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SCAN":
			var outNode Scan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "INDEXSCAN":
			var outNode IndexScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "INDEXONLYSCAN":
			var outNode IndexOnlyScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "BITMAPINDEXSCAN":
			var outNode BitmapIndexScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "BITMAPHEAPSCAN":
			var outNode BitmapHeapScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "TIDSCAN":
			var outNode TidScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SUBQUERYSCAN":
			var outNode SubqueryScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FUNCTIONSCAN":
			var outNode FunctionScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "VALUESSCAN":
			var outNode ValuesScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CTESCAN":
			var outNode CteScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "WORKTABLESCAN":
			var outNode WorkTableScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FOREIGNSCAN":
			var outNode ForeignScan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "JOIN":
			var outNode Join
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "NESTLOOP":
			var outNode NestLoop
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "NESTLOOPPARAM":
			var outNode NestLoopParam
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MERGEJOIN":
			var outNode MergeJoin
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "HASHJOIN":
			var outNode HashJoin
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MATERIAL":
			var outNode Material
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SORT":
			var outNode Sort
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "GROUP":
			var outNode Group
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "AGG":
			var outNode Agg
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "WINDOWAGG":
			var outNode WindowAgg
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "UNIQUE":
			var outNode Unique
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "HASH":
			var outNode Hash
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SETOP":
			var outNode SetOp
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "LOCKROWS":
			var outNode LockRows
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "LIMIT":
			var outNode Limit
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PLANROWMARK":
			var outNode PlanRowMark
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PLANINVALITEM":
			var outNode PlanInvalItem
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALIAS":
			var outNode Alias
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RANGEVAR":
			var outNode RangeVar
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "INTOCLAUSE":
			var outNode IntoClause
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "EXPR":
			var outNode Expr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "VAR":
			var outNode Var
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CONST":
			var outNode Const
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PARAM":
			var outNode Param
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "AGGREF":
			var outNode Aggref
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "WINDOWFUNC":
			var outNode WindowFunc
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ARRAYREF":
			var outNode ArrayRef
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FUNCEXPR":
			var outNode FuncExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "NAMEDARGEXPR":
			var outNode NamedArgExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "OPEXPR":
			var outNode OpExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SCALARARRAYOPEXPR":
			var outNode ScalarArrayOpExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "BOOLEXPR":
			var outNode BoolExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SUBLINK":
			var outNode SubLink
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SUBPLAN":
			var outNode SubPlan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ALTERNATIVESUBPLAN":
			var outNode AlternativeSubPlan
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FIELDSELECT":
			var outNode FieldSelect
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FIELDSTORE":
			var outNode FieldStore
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RELABELTYPE":
			var outNode RelabelType
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COERCEVIAIO":
			var outNode CoerceViaIO
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ARRAYCOERCEEXPR":
			var outNode ArrayCoerceExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CONVERTROWTYPEEXPR":
			var outNode ConvertRowtypeExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COLLATE":
			var outNode CollateExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CASE":
			var outNode CaseExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "WHEN":
			var outNode CaseWhen
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CASETESTEXPR":
			var outNode CaseTestExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ARRAY":
			var outNode ArrayExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ROW":
			var outNode RowExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "ROWCOMPARE":
			var outNode RowCompareExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COALESCE":
			var outNode CoalesceExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MINMAX":
			var outNode MinMaxExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "XMLEXPR":
			var outNode XmlExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COERCETODOMAIN":
			var outNode CoerceToDomain
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "COERCETODOMAINVALUE":
			var outNode CoerceToDomainValue
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SETTODEFAULT":
			var outNode SetToDefault
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "CURRENTOFEXPR":
			var outNode CurrentOfExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "TARGETENTRY":
			var outNode TargetEntry
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RANGETBLREF":
			var outNode RangeTblRef
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "JOINEXPR":
			var outNode JoinExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FROMEXPR":
			var outNode FromExpr
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "QUALCOST":
			var outNode QualCost
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "AGGCLAUSECOSTS":
			var outNode AggClauseCosts
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PLANNERGLOBAL":
			var outNode PlannerGlobal
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PLANNERINFO":
			var outNode PlannerInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RELOPTINFO":
			var outNode RelOptInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "INDEXOPTINFO":
			var outNode IndexOptInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "EQUIVALENCECLASS":
			var outNode EquivalenceClass
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "EQUIVALENCEMEMBER":
			var outNode EquivalenceMember
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PATHKEY":
			var outNode PathKey
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PARAMPATHINFO":
			var outNode ParamPathInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PATH":
			var outNode Path
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "INDEXPATH":
			var outNode IndexPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "BITMAPHEAPPATH":
			var outNode BitmapHeapPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "BITMAPANDPATH":
			var outNode BitmapAndPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "BITMAPORPATH":
			var outNode BitmapOrPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "TIDPATH":
			var outNode TidPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "FOREIGNPATH":
			var outNode ForeignPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "APPENDPATH":
			var outNode AppendPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MERGEAPPENDPATH":
			var outNode MergeAppendPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RESULTPATH":
			var outNode ResultPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MATERIALPATH":
			var outNode MaterialPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "UNIQUEPATH":
			var outNode UniquePath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "JOINPATH":
			var outNode JoinPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MERGEPATH":
			var outNode MergePath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "HASHPATH":
			var outNode HashPath
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "RESTRICTINFO":
			var outNode RestrictInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MERGESCANSELCACHE":
			var outNode MergeScanSelCache
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PLACEHOLDERVAR":
			var outNode PlaceHolderVar
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SPECIALJOININFO":
			var outNode SpecialJoinInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "LATERALJOININFO":
			var outNode LateralJoinInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "APPENDRELINFO":
			var outNode AppendRelInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PLACEHOLDERINFO":
			var outNode PlaceHolderInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "MINMAXAGGINFO":
			var outNode MinMaxAggInfo
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PLANNERPARAMITEM":
			var outNode PlannerParamItem
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "SEMIANTIJOINFACTORS":
			var outNode SemiAntiJoinFactors
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "JOINCOSTWORKSPACE":
			var outNode JoinCostWorkspace
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PARAMEXTERNDATA":
			var outNode ParamExternData
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PARAMLISTINFODATA":
			var outNode ParamListInfoData
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "PARAMEXECDATA":
			var outNode ParamExecData
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "VARATT_EXTERNAL":
			var outNode varatt_external
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "BLOCKIDDATA":
			var outNode BlockIdData
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		case "A_CONST":
			var outNode A_Const
			err = json.Unmarshal(jsonText, &outNode)
			if err != nil {
				return
			}
			node = outNode
		}
	}

	return
}
