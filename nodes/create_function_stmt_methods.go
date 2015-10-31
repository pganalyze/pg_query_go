// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateFunctionStmt) MarshalJSON() ([]byte, error) {
	type CreateFunctionStmtMarshalAlias CreateFunctionStmt
	return json.Marshal(map[string]interface{}{
		"CREATEFUNCTIONSTMT": (*CreateFunctionStmtMarshalAlias)(&node),
	})
}

func (node *CreateFunctionStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateFunctionStmt) Deparse() string {
	panic("Not Implemented")
}
