// Auto-generated - DO NOT EDIT

package pg_query

type BoolExpr struct {
	Xpr      Expr         `json:"xpr"`
	Boolop   BoolExprType `json:"boolop"`
	Args     []Node       `json:"args"`     /* arguments to this expression */
	Location int          `json:"location"` /* token location, or -1 if unknown */
}
