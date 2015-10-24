// Auto-generated - DO NOT EDIT

package pg_query

type CoerceViaIO struct {
	Xpr        Expr  `json:"xpr"`
	Arg        *Expr `json:"arg"`        /* input expression */
	Resulttype Oid   `json:"resulttype"` /* output type of coercion */
	/* output typmod is not stored, but is presumed -1 */
	Resultcollid Oid          `json:"resultcollid"` /* OID of collation, or InvalidOid if none */
	Coerceformat CoercionForm `json:"coerceformat"` /* how to display this node */
	Location     int          `json:"location"`     /* token location, or -1 if unknown */
}
