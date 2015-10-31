// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ReindexStmt) MarshalJSON() ([]byte, error) {
	type ReindexStmtMarshalAlias ReindexStmt
	return json.Marshal(map[string]interface{}{
		"REINDEXSTMT": (*ReindexStmtMarshalAlias)(&node),
	})
}

func (node *ReindexStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ReindexStmt) Deparse() string {
	panic("Not Implemented")
}
