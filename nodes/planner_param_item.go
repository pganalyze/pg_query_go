// Auto-generated - DO NOT EDIT

package pg_query

type PlannerParamItem struct {
	Item    Node `json:"item"`    /* the Var, PlaceHolderVar, or Aggref */
	ParamId int  `json:"paramId"` /* its assigned PARAM_EXEC slot number */
}
