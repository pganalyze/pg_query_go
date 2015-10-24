// Auto-generated - DO NOT EDIT

package pg_query

type RelabelType struct {
	Xpr           Expr         `json:"xpr"`
	Arg           *Expr        `json:"arg"`           /* input expression */
	Resulttype    Oid          `json:"resulttype"`    /* output type of coercion expression */
	Resulttypmod  int32        `json:"resulttypmod"`  /* output typmod (usually -1) */
	Resultcollid  Oid          `json:"resultcollid"`  /* OID of collation, or InvalidOid if none */
	Relabelformat CoercionForm `json:"relabelformat"` /* how to display this node */
	Location      int          `json:"location"`      /* token location, or -1 if unknown */
}
