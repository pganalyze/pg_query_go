// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type PlannerParamItem struct {
	Item    Node `json:"item"`    /* the Var, PlaceHolderVar, or Aggref */
	ParamId int  `json:"paramId"` /* its assigned PARAM_EXEC slot number */
}

func (node PlannerParamItem) MarshalJSON() ([]byte, error) {
	type PlannerParamItemMarshalAlias PlannerParamItem
	return json.Marshal(map[string]interface{}{
		"PLANNERPARAMITEM": (*PlannerParamItemMarshalAlias)(&node),
	})
}

func (node *PlannerParamItem) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
