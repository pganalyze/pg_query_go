// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PlanRowMark) MarshalJSON() ([]byte, error) {
	type PlanRowMarkMarshalAlias PlanRowMark
	return json.Marshal(map[string]interface{}{
		"PLANROWMARK": (*PlanRowMarkMarshalAlias)(&node),
	})
}

func (node *PlanRowMark) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PlanRowMark) Deparse() string {
	panic("Not Implemented")
}
