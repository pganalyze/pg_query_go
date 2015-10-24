// Auto-generated - DO NOT EDIT

package pg_query

type CurrentOfExpr struct {
	Xpr         Expr    `json:"xpr"`
	Cvarno      Index   `json:"cvarno"`       /* RT index of target relation */
	CursorName  *string `json:"cursor_name"`  /* name of referenced cursor, or NULL */
	CursorParam int     `json:"cursor_param"` /* refcursor parameter number, or 0 */
}
