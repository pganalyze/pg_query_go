// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type A_Const struct {
	Val      Value `json:"val"`      /* value (includes type info, see value.h) */
	Location int   `json:"location"` /* token location, or -1 if unknown */
}

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
