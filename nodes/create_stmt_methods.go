// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateStmt) MarshalJSON() ([]byte, error) {
	type CreateStmtMarshalAlias CreateStmt
	return json.Marshal(map[string]interface{}{
		"CREATESTMT": (*CreateStmtMarshalAlias)(&node),
	})
}

func (node *CreateStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateStmt) Deparse() string {
	panic("Not Implemented")
}
