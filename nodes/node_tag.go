package pg_query

import "strings"

// TODO(LukasFittl): I don't think we actually need this

type NodeTag uint

/*
 * The first field of every node is NodeTag. Each node created (with makeNode)
 * will have one of the following tags as the value of its first field.
 *
 * Note that the numbers of the node tags are not contiguous. We left holes
 * here so that we can add more tags without changing the existing enum's.
 * (Since node tag numbers never exist outside backend memory, there's no
 * real harm in renumbering, it just costs a full rebuild ...)
 */
const (
	T_Invalid NodeTag = iota

	/*
	 * TAGS FOR EXECUTOR NODES (execnodes.h)
	 */
)
const (
	T_IndexInfo NodeTag = iota + 10
	T_ExprContext
	T_ProjectionInfo
	T_JunkFilter
	T_ResultRelInfo
	T_EState
	T_TupleTableSlot

	/*
	 * TAGS FOR PLAN NODES (plannodes.h)
	 */
)
const (
	T_Plan NodeTag = iota + 100
	T_Result
	T_ModifyTable
	T_Append
	T_MergeAppend
	T_RecursiveUnion
	T_BitmapAnd
	T_BitmapOr
	T_Scan
	T_SeqScan
	T_IndexScan
	T_IndexOnlyScan
	T_BitmapIndexScan
	T_BitmapHeapScan
	T_TidScan
	T_SubqueryScan
	T_FunctionScan
	T_ValuesScan
	T_CteScan
	T_WorkTableScan
	T_ForeignScan
	T_Join
	T_NestLoop
	T_MergeJoin
	T_HashJoin
	T_Material
	T_Sort
	T_Group
	T_Agg
	T_WindowAgg
	T_Unique
	T_Hash
	T_SetOp
	T_LockRows
	T_Limit
	/* these aren't subclasses of Plan: */
	T_NestLoopParam
	T_PlanRowMark
	T_PlanInvalItem

	/*
	 * TAGS FOR PLAN STATE NODES (execnodes.h)
	 *
	 * These should correspond one-to-one with Plan node types.
	 */
)
const (
	T_PlanState NodeTag = iota + 200
	T_ResultState
	T_ModifyTableState
	T_AppendState
	T_MergeAppendState
	T_RecursiveUnionState
	T_BitmapAndState
	T_BitmapOrState
	T_ScanState
	T_SeqScanState
	T_IndexScanState
	T_IndexOnlyScanState
	T_BitmapIndexScanState
	T_BitmapHeapScanState
	T_TidScanState
	T_SubqueryScanState
	T_FunctionScanState
	T_ValuesScanState
	T_CteScanState
	T_WorkTableScanState
	T_ForeignScanState
	T_JoinState
	T_NestLoopState
	T_MergeJoinState
	T_HashJoinState
	T_MaterialState
	T_SortState
	T_GroupState
	T_AggState
	T_WindowAggState
	T_UniqueState
	T_HashState
	T_SetOpState
	T_LockRowsState
	T_LimitState

	/*
	 * TAGS FOR PRIMITIVE NODES (primnodes.h)
	 */
)
const (
	T_Alias NodeTag = iota + 300
	T_RangeVar
	T_Expr
	T_Var
	T_Const
	T_Param
	T_Aggref
	T_WindowFunc
	T_ArrayRef
	T_FuncExpr
	T_NamedArgExpr
	T_OpExpr
	T_DistinctExpr
	T_NullIfExpr
	T_ScalarArrayOpExpr
	T_BoolExpr
	T_SubLink
	T_SubPlan
	T_AlternativeSubPlan
	T_FieldSelect
	T_FieldStore
	T_RelabelType
	T_CoerceViaIO
	T_ArrayCoerceExpr
	T_ConvertRowtypeExpr
	T_CollateExpr
	T_CaseExpr
	T_CaseWhen
	T_CaseTestExpr
	T_ArrayExpr
	T_RowExpr
	T_RowCompareExpr
	T_CoalesceExpr
	T_MinMaxExpr
	T_XmlExpr
	T_NullTest
	T_BooleanTest
	T_CoerceToDomain
	T_CoerceToDomainValue
	T_SetToDefault
	T_CurrentOfExpr
	T_TargetEntry
	T_RangeTblRef
	T_JoinExpr
	T_FromExpr
	T_IntoClause

	/*
	 * TAGS FOR EXPRESSION STATE NODES (execnodes.h)
	 *
	 * These correspond (not always one-for-one) to primitive nodes derived
	 * from Expr.
	 */
)
const (
	T_ExprState NodeTag = iota + 400
	T_GenericExprState
	T_WholeRowVarExprState
	T_AggrefExprState
	T_WindowFuncExprState
	T_ArrayRefExprState
	T_FuncExprState
	T_ScalarArrayOpExprState
	T_BoolExprState
	T_SubPlanState
	T_AlternativeSubPlanState
	T_FieldSelectState
	T_FieldStoreState
	T_CoerceViaIOState
	T_ArrayCoerceExprState
	T_ConvertRowtypeExprState
	T_CaseExprState
	T_CaseWhenState
	T_ArrayExprState
	T_RowExprState
	T_RowCompareExprState
	T_CoalesceExprState
	T_MinMaxExprState
	T_XmlExprState
	T_NullTestState
	T_CoerceToDomainState
	T_DomainConstraintState

	/*
	 * TAGS FOR PLANNER NODES (relation.h)
	 */
)
const (
	T_PlannerInfo NodeTag = iota + 500
	T_PlannerGlobal
	T_RelOptInfo
	T_IndexOptInfo
	T_ParamPathInfo
	T_Path
	T_IndexPath
	T_BitmapHeapPath
	T_BitmapAndPath
	T_BitmapOrPath
	T_NestPath
	T_MergePath
	T_HashPath
	T_TidPath
	T_ForeignPath
	T_AppendPath
	T_MergeAppendPath
	T_ResultPath
	T_MaterialPath
	T_UniquePath
	T_EquivalenceClass
	T_EquivalenceMember
	T_PathKey
	T_RestrictInfo
	T_PlaceHolderVar
	T_SpecialJoinInfo
	T_LateralJoinInfo
	T_AppendRelInfo
	T_PlaceHolderInfo
	T_MinMaxAggInfo
	T_PlannerParamItem

	/*
	 * TAGS FOR MEMORY NODES (memnodes.h)
	 */
)
const (
	T_MemoryContext NodeTag = iota + 600
	T_AllocSetContext

	/*
	 * TAGS FOR VALUE NODES (value.h)
	 */
)
const (
	T_Value NodeTag = iota + 650
	T_Integer
	T_Float
	T_String
	T_BitString
	T_Null

	/*
	 * TAGS FOR LIST NODES (pg_list.h)
	 */
	T_List
	T_IntList
	T_OidList

	/*
	 * TAGS FOR STATEMENT NODES (mostly in parsenodes.h)
	 */
)
const (
	T_Query NodeTag = iota + 700
	T_PlannedStmt
	T_InsertStmt
	T_DeleteStmt
	T_UpdateStmt
	T_SelectStmt
	T_AlterTableStmt
	T_AlterTableCmd
	T_AlterDomainStmt
	T_SetOperationStmt
	T_GrantStmt
	T_GrantRoleStmt
	T_AlterDefaultPrivilegesStmt
	T_ClosePortalStmt
	T_ClusterStmt
	T_CopyStmt
	T_CreateStmt
	T_DefineStmt
	T_DropStmt
	T_TruncateStmt
	T_CommentStmt
	T_FetchStmt
	T_IndexStmt
	T_CreateFunctionStmt
	T_AlterFunctionStmt
	T_DoStmt
	T_RenameStmt
	T_RuleStmt
	T_NotifyStmt
	T_ListenStmt
	T_UnlistenStmt
	T_TransactionStmt
	T_ViewStmt
	T_LoadStmt
	T_CreateDomainStmt
	T_CreatedbStmt
	T_DropdbStmt
	T_VacuumStmt
	T_ExplainStmt
	T_CreateTableAsStmt
	T_CreateSeqStmt
	T_AlterSeqStmt
	T_VariableSetStmt
	T_VariableShowStmt
	T_DiscardStmt
	T_CreateTrigStmt
	T_CreatePLangStmt
	T_CreateRoleStmt
	T_AlterRoleStmt
	T_DropRoleStmt
	T_LockStmt
	T_ConstraintsSetStmt
	T_ReindexStmt
	T_CheckPointStmt
	T_CreateSchemaStmt
	T_AlterDatabaseStmt
	T_AlterDatabaseSetStmt
	T_AlterRoleSetStmt
	T_CreateConversionStmt
	T_CreateCastStmt
	T_CreateOpClassStmt
	T_CreateOpFamilyStmt
	T_AlterOpFamilyStmt
	T_PrepareStmt
	T_ExecuteStmt
	T_DeallocateStmt
	T_DeclareCursorStmt
	T_CreateTableSpaceStmt
	T_DropTableSpaceStmt
	T_AlterObjectSchemaStmt
	T_AlterOwnerStmt
	T_DropOwnedStmt
	T_ReassignOwnedStmt
	T_CompositeTypeStmt
	T_CreateEnumStmt
	T_CreateRangeStmt
	T_AlterEnumStmt
	T_AlterTSDictionaryStmt
	T_AlterTSConfigurationStmt
	T_CreateFdwStmt
	T_AlterFdwStmt
	T_CreateForeignServerStmt
	T_AlterForeignServerStmt
	T_CreateUserMappingStmt
	T_AlterUserMappingStmt
	T_DropUserMappingStmt
	T_AlterTableSpaceOptionsStmt
	T_AlterTableMoveAllStmt
	T_SecLabelStmt
	T_CreateForeignTableStmt
	T_CreateExtensionStmt
	T_AlterExtensionStmt
	T_AlterExtensionContentsStmt
	T_CreateEventTrigStmt
	T_AlterEventTrigStmt
	T_RefreshMatViewStmt
	T_ReplicaIdentityStmt
	T_AlterSystemStmt

	/*
	 * TAGS FOR PARSE TREE NODES (parsenodes.h)
	 */
)
const (
	T_A_Expr NodeTag = iota + 900
	T_ColumnRef
	T_ParamRef
	T_A_Const
	T_FuncCall
	T_A_Star
	T_A_Indices
	T_A_Indirection
	T_A_ArrayExpr
	T_ResTarget
	T_TypeCast
	T_CollateClause
	T_SortBy
	T_WindowDef
	T_RangeSubselect
	T_RangeFunction
	T_TypeName
	T_ColumnDef
	T_IndexElem
	T_Constraint
	T_DefElem
	T_RangeTblEntry
	T_RangeTblFunction
	T_WithCheckOption
	T_SortGroupClause
	T_WindowClause
	T_PrivGrantee
	T_FuncWithArgs
	T_AccessPriv
	T_CreateOpClassItem
	T_TableLikeClause
	T_FunctionParameter
	T_LockingClause
	T_RowMarkClause
	T_XmlSerialize
	T_WithClause
	T_CommonTableExpr

	/*
	 * TAGS FOR REPLICATION GRAMMAR PARSE NODES (replnodes.h)
	 */
	T_IdentifySystemCmd
	T_BaseBackupCmd
	T_CreateReplicationSlotCmd
	T_DropReplicationSlotCmd
	T_StartReplicationCmd
	T_TimeLineHistoryCmd

	/*
	 * TAGS FOR RANDOM OTHER STUFF
	 *
	 * These are objects that aren't part of parse/plan/execute node tree
	 * structures, but we give them NodeTags anyway for identification
	 * purposes (usually because they are involved in APIs where we want to
	 * pass multiple object types through the same pointer).
	 */
)
const (
	T_TriggerData      NodeTag = iota + 950 /* in commands/trigger.h */
	T_EventTriggerData                      /* in commands/event_trigger.h */
	T_ReturnSetInfo                         /* in nodes/execnodes.h */
	T_WindowObjectData                      /* private in nodeWindowAgg.c */
	T_TIDBitmap                             /* in nodes/tidbitmap.h */
	T_InlineCodeBlock                       /* in nodes/parsenodes.h */
	T_FdwRoutine                            /* in foreign/fdwapi.h */
)

func JsonKeyToNodeTag(input string) NodeTag {
	reversed_map := make(map[string]NodeTag)

	if input == "SELECT" {
		input = "SELECTSTMT"
	}

	for key, value := range _NodeTag_map {
		reversed_map[strings.ToUpper(strings.TrimPrefix(value, "T_"))] = key
	}

	return reversed_map[input]
}
