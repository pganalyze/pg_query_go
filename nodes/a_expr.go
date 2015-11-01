// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type A_Expr struct {
	Kind     A_Expr_Kind `json:"kind"`     /* see above */
	Name     []Node      `json:"name"`     /* possibly-qualified name of operator */
	Lexpr    Node        `json:"lexpr"`    /* left argument, or NULL if none */
	Rexpr    Node        `json:"rexpr"`    /* right argument, or NULL if none */
	Location int         `json:"location"` /* token location, or -1 if unknown */
}

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
