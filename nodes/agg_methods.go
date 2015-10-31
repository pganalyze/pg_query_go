// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Agg) MarshalJSON() ([]byte, error) {
	type AggMarshalAlias Agg
	return json.Marshal(map[string]interface{}{
		"AGG": (*AggMarshalAlias)(&node),
	})
}

func (node *Agg) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Agg) Deparse() string {
	panic("Not Implemented")
}
