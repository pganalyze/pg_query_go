// Auto-generated - DO NOT EDIT

package pg_query

type CaseExpr struct {
	Xpr        Expr   `json:"xpr"`
	Casetype   Oid    `json:"casetype"`   /* type of expression result */
	Casecollid Oid    `json:"casecollid"` /* OID of collation, or InvalidOid if none */
	Arg        *Expr  `json:"arg"`        /* implicit equality comparison argument */
	Args       []Node `json:"args"`       /* the arguments (list of WHEN clauses) */
	Defresult  *Expr  `json:"defresult"`  /* the default result (ELSE clause) */
	Location   int    `json:"location"`   /* token location, or -1 if unknown */
}
