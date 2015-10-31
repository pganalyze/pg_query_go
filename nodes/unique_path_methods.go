// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node UniquePath) MarshalJSON() ([]byte, error) {
	type UniquePathMarshalAlias UniquePath
	return json.Marshal(map[string]interface{}{
		"UNIQUEPATH": (*UniquePathMarshalAlias)(&node),
	})
}

func (node *UniquePath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node UniquePath) Deparse() string {
	panic("Not Implemented")
}
