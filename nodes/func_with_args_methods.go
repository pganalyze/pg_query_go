// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node FuncWithArgs) MarshalJSON() ([]byte, error) {
	type FuncWithArgsMarshalAlias FuncWithArgs
	return json.Marshal(map[string]interface{}{
		"FUNCWITHARGS": (*FuncWithArgsMarshalAlias)(&node),
	})
}

func (node *FuncWithArgs) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node FuncWithArgs) Deparse() string {
	panic("Not Implemented")
}
