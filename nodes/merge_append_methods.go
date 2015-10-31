// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node MergeAppend) MarshalJSON() ([]byte, error) {
	type MergeAppendMarshalAlias MergeAppend
	return json.Marshal(map[string]interface{}{
		"MERGEAPPEND": (*MergeAppendMarshalAlias)(&node),
	})
}

func (node *MergeAppend) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node MergeAppend) Deparse() string {
	panic("Not Implemented")
}
