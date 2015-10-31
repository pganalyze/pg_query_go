// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ExecuteStmt) MarshalJSON() ([]byte, error) {
	type ExecuteStmtMarshalAlias ExecuteStmt
	return json.Marshal(map[string]interface{}{
		"EXECUTESTMT": (*ExecuteStmtMarshalAlias)(&node),
	})
}

func (node *ExecuteStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ExecuteStmt) Deparse() string {
	panic("Not Implemented")
}
