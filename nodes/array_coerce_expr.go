// Auto-generated - DO NOT EDIT

package pg_query

type ArrayCoerceExpr struct {
	Xpr          Expr         `json:"xpr"`
	Arg          *Expr        `json:"arg"`          /* input expression (yields an array) */
	Elemfuncid   Oid          `json:"elemfuncid"`   /* OID of element coercion function, or 0 */
	Resulttype   Oid          `json:"resulttype"`   /* output type of coercion (an array type) */
	Resulttypmod int32        `json:"resulttypmod"` /* output typmod (also element typmod) */
	Resultcollid Oid          `json:"resultcollid"` /* OID of collation, or InvalidOid if none */
	IsExplicit   bool         `json:"isExplicit"`   /* conversion semantics flag to pass to func */
	Coerceformat CoercionForm `json:"coerceformat"` /* how to display this node */
	Location     int          `json:"location"`     /* token location, or -1 if unknown */
}
