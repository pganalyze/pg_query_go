// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ResultPath) MarshalJSON() ([]byte, error) {
	type ResultPathMarshalAlias ResultPath
	return json.Marshal(map[string]interface{}{
		"RESULTPATH": (*ResultPathMarshalAlias)(&node),
	})
}

func (node *ResultPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ResultPath) Deparse() string {
	panic("Not Implemented")
}
