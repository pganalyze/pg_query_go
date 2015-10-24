// Auto-generated - DO NOT EDIT

package pg_query

type ConvertRowtypeExpr struct {
	Xpr        Expr  `json:"xpr"`
	Arg        *Expr `json:"arg"`        /* input expression */
	Resulttype Oid   `json:"resulttype"` /* output type (always a composite type) */
	/* Like RowExpr, we deliberately omit a typmod and collation here */
	Convertformat CoercionForm `json:"convertformat"` /* how to display this node */
	Location      int          `json:"location"`      /* token location, or -1 if unknown */
}
