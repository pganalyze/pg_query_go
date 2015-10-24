// Auto-generated - DO NOT EDIT

package pg_query

type FieldSelect struct {
	Xpr        Expr       `json:"xpr"`
	Arg        *Expr      `json:"arg"`        /* input expression */
	Fieldnum   AttrNumber `json:"fieldnum"`   /* attribute number of field to extract */
	Resulttype Oid        `json:"resulttype"` /* type of the field (result type of this
	 * node) */
	Resulttypmod int32 `json:"resulttypmod"` /* output typmod (usually -1) */
	Resultcollid Oid   `json:"resultcollid"` /* OID of collation of the field */
}
