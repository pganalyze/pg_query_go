// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node A_Const) MarshalJSON() ([]byte, error) {
	type A_ConstMarshalAlias A_Const
	return json.Marshal(map[string]interface{}{
		"A_CONST": (*A_ConstMarshalAlias)(&node),
	})
}

func (node *A_Const) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node A_Const) Deparse() string {
	panic("Not Implemented")
}
