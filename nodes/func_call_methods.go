// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node FuncCall) MarshalJSON() ([]byte, error) {
	type FuncCallMarshalAlias FuncCall
	return json.Marshal(map[string]interface{}{
		"FUNCCALL": (*FuncCallMarshalAlias)(&node),
	})
}

func (node *FuncCall) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node FuncCall) Deparse() string {
	panic("Not Implemented")
}
