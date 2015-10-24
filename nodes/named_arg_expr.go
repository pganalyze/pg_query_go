// Auto-generated - DO NOT EDIT

package pg_query

type NamedArgExpr struct {
	Xpr       Expr    `json:"xpr"`
	Arg       *Expr   `json:"arg"`       /* the argument expression */
	Name      *string `json:"name"`      /* the name */
	Argnumber int     `json:"argnumber"` /* argument's number in positional notation */
	Location  int     `json:"location"`  /* argument name location, or -1 if unknown */
}
