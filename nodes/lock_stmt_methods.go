// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node LockStmt) Deparse() string {
	panic("Not Implemented")
}
