// Auto-generated - DO NOT EDIT

package pg_query

type CoalesceExpr struct {
	Xpr            Expr   `json:"xpr"`
	Coalescetype   Oid    `json:"coalescetype"`   /* type of expression result */
	Coalescecollid Oid    `json:"coalescecollid"` /* OID of collation, or InvalidOid if none */
	Args           []Node `json:"args"`           /* the arguments */
	Location       int    `json:"location"`       /* token location, or -1 if unknown */
}
