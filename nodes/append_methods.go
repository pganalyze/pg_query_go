// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Append) MarshalJSON() ([]byte, error) {
	type AppendMarshalAlias Append
	return json.Marshal(map[string]interface{}{
		"APPEND": (*AppendMarshalAlias)(&node),
	})
}

func (node *Append) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Append) Deparse() string {
	panic("Not Implemented")
}
