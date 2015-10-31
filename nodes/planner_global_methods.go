// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PlannerGlobal) MarshalJSON() ([]byte, error) {
	type PlannerGlobalMarshalAlias PlannerGlobal
	return json.Marshal(map[string]interface{}{
		"PLANNERGLOBAL": (*PlannerGlobalMarshalAlias)(&node),
	})
}

func (node *PlannerGlobal) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PlannerGlobal) Deparse() string {
	panic("Not Implemented")
}
