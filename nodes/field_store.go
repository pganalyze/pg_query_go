// Auto-generated - DO NOT EDIT

package pg_query

type FieldStore struct {
	Xpr        Expr   `json:"xpr"`
	Arg        *Expr  `json:"arg"`        /* input tuple value */
	Newvals    []Node `json:"newvals"`    /* new value(s) for field(s) */
	Fieldnums  []Node `json:"fieldnums"`  /* integer list of field attnums */
	Resulttype Oid    `json:"resulttype"` /* type of result (same as type of arg) */
	/* Like RowExpr, we deliberately omit a typmod and collation here */
}
