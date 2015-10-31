// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node MergePath) MarshalJSON() ([]byte, error) {
	type MergePathMarshalAlias MergePath
	return json.Marshal(map[string]interface{}{
		"MERGEPATH": (*MergePathMarshalAlias)(&node),
	})
}

func (node *MergePath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node MergePath) Deparse() string {
	panic("Not Implemented")
}
