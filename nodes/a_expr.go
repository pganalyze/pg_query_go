// Auto-generated - DO NOT EDIT

package pg_query

type A_Expr struct {
	Kind     A_Expr_Kind `json:"kind"`     /* see above */
	Name     []Node      `json:"name"`     /* possibly-qualified name of operator */
	Lexpr    Node        `json:"lexpr"`    /* left argument, or NULL if none */
	Rexpr    Node        `json:"rexpr"`    /* right argument, or NULL if none */
	Location int         `json:"location"` /* token location, or -1 if unknown */
}
