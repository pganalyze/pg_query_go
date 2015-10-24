// Auto-generated - DO NOT EDIT

package pg_query

type RowExpr struct {
	Xpr       Expr   `json:"xpr"`
	Args      []Node `json:"args"`       /* the fields */
	RowTypeid Oid    `json:"row_typeid"` /* RECORDOID or a composite type's ID */

	/*
	 * Note: we deliberately do NOT store a typmod.  Although a typmod will be
	 * associated with specific RECORD types at runtime, it will differ for
	 * different backends, and so cannot safely be stored in stored
	 * parsetrees.  We must assume typmod -1 for a RowExpr node.
	 *
	 * We don't need to store a collation either.  The result type is
	 * necessarily composite, and composite types never have a collation.
	 */
	RowFormat CoercionForm `json:"row_format"` /* how to display this node */
	Colnames  []Node       `json:"colnames"`   /* list of String, or NIL */
	Location  int          `json:"location"`   /* token location, or -1 if unknown */
}
