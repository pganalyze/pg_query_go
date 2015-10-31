// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DropdbStmt) MarshalJSON() ([]byte, error) {
	type DropdbStmtMarshalAlias DropdbStmt
	return json.Marshal(map[string]interface{}{
		"DROPDBSTMT": (*DropdbStmtMarshalAlias)(&node),
	})
}

func (node *DropdbStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DropdbStmt) Deparse() string {
	panic("Not Implemented")
}
