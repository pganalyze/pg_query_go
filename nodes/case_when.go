// Auto-generated - DO NOT EDIT

package pg_query

type CaseWhen struct {
	Xpr      Expr  `json:"xpr"`
	Expr     *Expr `json:"expr"`     /* condition expression */
	Result   *Expr `json:"result"`   /* substitution result */
	Location int   `json:"location"` /* token location, or -1 if unknown */
}
