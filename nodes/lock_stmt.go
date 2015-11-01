// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type LockStmt struct {
	Relations []Node `json:"relations"` /* relations to lock */
	Mode      int    `json:"mode"`      /* lock mode */
	Nowait    bool   `json:"nowait"`    /* no wait mode */
}

func (node LockStmt) MarshalJSON() ([]byte, error) {
	type LockStmtMarshalAlias LockStmt
	return json.Marshal(map[string]interface{}{
		"LOCK": (*LockStmtMarshalAlias)(&node),
	})
}

func (node *LockStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
