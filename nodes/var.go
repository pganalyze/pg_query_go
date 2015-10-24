// Auto-generated - DO NOT EDIT

package pg_query

type Var struct {
	Xpr   Expr  `json:"xpr"`
	Varno Index `json:"varno"` /* index of this var's relation in the range
	 * table, or INNER_VAR/OUTER_VAR/INDEX_VAR */
	Varattno AttrNumber `json:"varattno"` /* attribute number of this var, or zero for
	 * all */
	Vartype     Oid   `json:"vartype"`     /* pg_type OID for the type of this var */
	Vartypmod   int32 `json:"vartypmod"`   /* pg_attribute typmod value */
	Varcollid   Oid   `json:"varcollid"`   /* OID of collation, or InvalidOid if none */
	Varlevelsup Index `json:"varlevelsup"` /* for subquery variables referencing outer
	 * relations; 0 in a normal var, >0 means N
	 * levels up */
	Varnoold  Index      `json:"varnoold"`  /* original value of varno, for debugging */
	Varoattno AttrNumber `json:"varoattno"` /* original value of varattno */
	Location  int        `json:"location"`  /* token location, or -1 if unknown */
}
