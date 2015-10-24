// Auto-generated - DO NOT EDIT

package pg_query

type LockStmt struct {
	Relations []Node `json:"relations"` /* relations to lock */
	Mode      int    `json:"mode"`      /* lock mode */
	Nowait    bool   `json:"nowait"`    /* no wait mode */
}
