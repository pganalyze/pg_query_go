// Auto-generated - DO NOT EDIT

package pg_query

type PlannerGlobal struct {
	BoundParams ParamListInfo `json:"boundParams"` /* Param values provided to planner() */

	Subplans []Node `json:"subplans"` /* Plans for SubPlan nodes */

	Subroots []Node `json:"subroots"` /* PlannerInfos for SubPlan nodes */

	RewindPlanIds []uint32 `json:"rewindPlanIDs"` /* indices of subplans that require REWIND */

	Finalrtable []Node `json:"finalrtable"` /* "flat" rangetable for executor */

	Finalrowmarks []Node `json:"finalrowmarks"` /* "flat" list of PlanRowMarks */

	ResultRelations []Node `json:"resultRelations"` /* "flat" list of integer RT indexes */

	RelationOids []Node `json:"relationOids"` /* OIDs of relations the plan depends on */

	InvalItems []Node `json:"invalItems"` /* other dependencies, as PlanInvalItems */

	NParamExec int `json:"nParamExec"` /* number of PARAM_EXEC Params used */

	LastPhid Index `json:"lastPHId"` /* highest PlaceHolderVar ID assigned */

	LastRowMarkId Index `json:"lastRowMarkId"` /* highest PlanRowMark ID assigned */

	TransientPlan bool `json:"transientPlan"` /* redo plan when TransactionXmin changes? */
}
