// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node TidPath) MarshalJSON() ([]byte, error) {
	type TidPathMarshalAlias TidPath
	return json.Marshal(map[string]interface{}{
		"TIDPATH": (*TidPathMarshalAlias)(&node),
	})
}

func (node *TidPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node TidPath) Deparse() string {
	panic("Not Implemented")
}
