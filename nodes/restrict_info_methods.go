// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RestrictInfo) MarshalJSON() ([]byte, error) {
	type RestrictInfoMarshalAlias RestrictInfo
	return json.Marshal(map[string]interface{}{
		"RESTRICTINFO": (*RestrictInfoMarshalAlias)(&node),
	})
}

func (node *RestrictInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RestrictInfo) Deparse() string {
	panic("Not Implemented")
}
