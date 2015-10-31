// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node HashPath) MarshalJSON() ([]byte, error) {
	type HashPathMarshalAlias HashPath
	return json.Marshal(map[string]interface{}{
		"HASHPATH": (*HashPathMarshalAlias)(&node),
	})
}

func (node *HashPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node HashPath) Deparse() string {
	panic("Not Implemented")
}
