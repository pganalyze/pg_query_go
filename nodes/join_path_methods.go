// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node JoinPath) MarshalJSON() ([]byte, error) {
	type JoinPathMarshalAlias JoinPath
	return json.Marshal(map[string]interface{}{
		"JOINPATH": (*JoinPathMarshalAlias)(&node),
	})
}

func (node *JoinPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node JoinPath) Deparse() string {
	panic("Not Implemented")
}
