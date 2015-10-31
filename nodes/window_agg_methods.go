// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node WindowAgg) MarshalJSON() ([]byte, error) {
	type WindowAggMarshalAlias WindowAgg
	return json.Marshal(map[string]interface{}{
		"WINDOWAGG": (*WindowAggMarshalAlias)(&node),
	})
}

func (node *WindowAgg) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node WindowAgg) Deparse() string {
	panic("Not Implemented")
}
