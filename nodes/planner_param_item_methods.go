// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node PlannerParamItem) Deparse() string {
	panic("Not Implemented")
}
