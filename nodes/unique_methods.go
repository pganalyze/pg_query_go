// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Unique) MarshalJSON() ([]byte, error) {
	type UniqueMarshalAlias Unique
	return json.Marshal(map[string]interface{}{
		"UNIQUE": (*UniqueMarshalAlias)(&node),
	})
}

func (node *Unique) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Unique) Deparse() string {
	panic("Not Implemented")
}
