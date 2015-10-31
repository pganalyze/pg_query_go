// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreatedbStmt) MarshalJSON() ([]byte, error) {
	type CreatedbStmtMarshalAlias CreatedbStmt
	return json.Marshal(map[string]interface{}{
		"CREATEDBSTMT": (*CreatedbStmtMarshalAlias)(&node),
	})
}

func (node *CreatedbStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreatedbStmt) Deparse() string {
	panic("Not Implemented")
}
