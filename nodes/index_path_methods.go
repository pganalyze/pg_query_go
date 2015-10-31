// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node IndexPath) MarshalJSON() ([]byte, error) {
	type IndexPathMarshalAlias IndexPath
	return json.Marshal(map[string]interface{}{
		"INDEXPATH": (*IndexPathMarshalAlias)(&node),
	})
}

func (node *IndexPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node IndexPath) Deparse() string {
	panic("Not Implemented")
}
