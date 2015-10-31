// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PlannerInfo) MarshalJSON() ([]byte, error) {
	type PlannerInfoMarshalAlias PlannerInfo
	return json.Marshal(map[string]interface{}{
		"PLANNERINFO": (*PlannerInfoMarshalAlias)(&node),
	})
}

func (node *PlannerInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PlannerInfo) Deparse() string {
	panic("Not Implemented")
}
