// Auto-generated - DO NOT EDIT

package pg_query

type CollateExpr struct {
	Xpr      Expr  `json:"xpr"`
	Arg      *Expr `json:"arg"`      /* input expression */
	CollOid  Oid   `json:"collOid"`  /* collation's OID */
	Location int   `json:"location"` /* token location, or -1 if unknown */
}
