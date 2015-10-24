// Auto-generated - DO NOT EDIT

package pg_query

type CoerceToDomain struct {
	Xpr            Expr         `json:"xpr"`
	Arg            *Expr        `json:"arg"`            /* input expression */
	Resulttype     Oid          `json:"resulttype"`     /* domain type ID (result type) */
	Resulttypmod   int32        `json:"resulttypmod"`   /* output typmod (currently always -1) */
	Resultcollid   Oid          `json:"resultcollid"`   /* OID of collation, or InvalidOid if none */
	Coercionformat CoercionForm `json:"coercionformat"` /* how to display this node */
	Location       int          `json:"location"`       /* token location, or -1 if unknown */
}
