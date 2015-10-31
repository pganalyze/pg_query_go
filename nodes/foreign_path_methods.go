// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ForeignPath) MarshalJSON() ([]byte, error) {
	type ForeignPathMarshalAlias ForeignPath
	return json.Marshal(map[string]interface{}{
		"FOREIGNPATH": (*ForeignPathMarshalAlias)(&node),
	})
}

func (node *ForeignPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ForeignPath) Deparse() string {
	panic("Not Implemented")
}
