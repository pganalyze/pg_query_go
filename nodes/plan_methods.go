// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Plan) MarshalJSON() ([]byte, error) {
	type PlanMarshalAlias Plan
	return json.Marshal(map[string]interface{}{
		"PLAN": (*PlanMarshalAlias)(&node),
	})
}

func (node *Plan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Plan) Deparse() string {
	panic("Not Implemented")
}
