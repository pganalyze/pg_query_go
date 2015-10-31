// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PlanInvalItem) MarshalJSON() ([]byte, error) {
	type PlanInvalItemMarshalAlias PlanInvalItem
	return json.Marshal(map[string]interface{}{
		"PLANINVALITEM": (*PlanInvalItemMarshalAlias)(&node),
	})
}

func (node *PlanInvalItem) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PlanInvalItem) Deparse() string {
	panic("Not Implemented")
}
