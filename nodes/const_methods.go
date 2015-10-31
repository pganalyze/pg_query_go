// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Const) MarshalJSON() ([]byte, error) {
	type ConstMarshalAlias Const
	return json.Marshal(map[string]interface{}{
		"CONST": (*ConstMarshalAlias)(&node),
	})
}

func (node *Const) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Const) Deparse() string {
	panic("Not Implemented")
}
