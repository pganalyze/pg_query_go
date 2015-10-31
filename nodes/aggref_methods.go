// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Aggref) MarshalJSON() ([]byte, error) {
	type AggrefMarshalAlias Aggref
	return json.Marshal(map[string]interface{}{
		"AGGREF": (*AggrefMarshalAlias)(&node),
	})
}

func (node *Aggref) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Aggref) Deparse() string {
	panic("Not Implemented")
}
