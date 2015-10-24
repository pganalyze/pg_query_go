// Auto-generated - DO NOT EDIT

package pg_query

type PlannedStmt struct {
	CommandType CmdType `json:"commandType"` /* select|insert|update|delete */

	QueryId uint32 `json:"queryId"` /* query identifier (copied from Query) */

	HasReturning bool `json:"hasReturning"` /* is it insert|update|delete RETURNING? */

	HasModifyingCte bool `json:"hasModifyingCTE"` /* has insert|update|delete in WITH? */

	CanSetTag bool `json:"canSetTag"` /* do I set the command result tag? */

	TransientPlan bool `json:"transientPlan"` /* redo plan when TransactionXmin changes? */

	PlanTree *Plan `json:"planTree"` /* tree of Plan nodes */

	Rtable []Node `json:"rtable"` /* list of RangeTblEntry nodes */

	/* rtable indexes of target relations for INSERT/UPDATE/DELETE */
	ResultRelations []Node `json:"resultRelations"` /* integer list of RT indexes, or NIL */

	UtilityStmt Node `json:"utilityStmt"` /* non-null if this is DECLARE CURSOR */

	Subplans []Node `json:"subplans"` /* Plan trees for SubPlan expressions */

	RewindPlanIds []uint32 `json:"rewindPlanIDs"` /* indices of subplans that require REWIND */

	RowMarks []Node `json:"rowMarks"` /* a list of PlanRowMark's */

	RelationOids []Node `json:"relationOids"` /* OIDs of relations the plan depends on */

	InvalItems []Node `json:"invalItems"` /* other dependencies, as PlanInvalItems */

	NParamExec int `json:"nParamExec"` /* number of PARAM_EXEC Params used */
}
