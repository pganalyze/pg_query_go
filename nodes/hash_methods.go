// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Hash) MarshalJSON() ([]byte, error) {
	type HashMarshalAlias Hash
	return json.Marshal(map[string]interface{}{
		"HASH": (*HashMarshalAlias)(&node),
	})
}

func (node *Hash) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Hash) Deparse() string {
	panic("Not Implemented")
}
