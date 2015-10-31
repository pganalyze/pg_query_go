// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node A_Expr) MarshalJSON() ([]byte, error) {
	type A_ExprMarshalAlias A_Expr
	return json.Marshal(map[string]interface{}{
		"AEXPR": (*A_ExprMarshalAlias)(&node),
	})
}

func (node *A_Expr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node A_Expr) Deparse() string {
	panic("Not Implemented")
}
