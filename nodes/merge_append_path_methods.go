// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node MergeAppendPath) MarshalJSON() ([]byte, error) {
	type MergeAppendPathMarshalAlias MergeAppendPath
	return json.Marshal(map[string]interface{}{
		"MERGEAPPENDPATH": (*MergeAppendPathMarshalAlias)(&node),
	})
}

func (node *MergeAppendPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node MergeAppendPath) Deparse() string {
	panic("Not Implemented")
}
