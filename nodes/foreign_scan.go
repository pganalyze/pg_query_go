// Auto-generated - DO NOT EDIT

package pg_query

type ForeignScan struct {
	Scan        Scan   `json:"scan"`
	FdwExprs    []Node `json:"fdw_exprs"`   /* expressions that FDW may evaluate */
	FdwPrivate  []Node `json:"fdw_private"` /* private data for FDW */
	FsSystemCol bool   `json:"fsSystemCol"` /* true if any "system column" is needed */
}
