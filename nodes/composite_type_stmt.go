// Auto-generated - DO NOT EDIT

package pg_query

type CompositeTypeStmt struct {
	Typevar    *RangeVar `json:"typevar"`    /* the composite type to be created */
	Coldeflist []Node    `json:"coldeflist"` /* list of ColumnDef nodes */
}
