// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateFdwStmt) MarshalJSON() ([]byte, error) {
	type CreateFdwStmtMarshalAlias CreateFdwStmt
	return json.Marshal(map[string]interface{}{
		"CREATEFDWSTMT": (*CreateFdwStmtMarshalAlias)(&node),
	})
}

func (node *CreateFdwStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateFdwStmt) Deparse() string {
	panic("Not Implemented")
}
