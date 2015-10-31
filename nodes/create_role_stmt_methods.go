// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateRoleStmt) MarshalJSON() ([]byte, error) {
	type CreateRoleStmtMarshalAlias CreateRoleStmt
	return json.Marshal(map[string]interface{}{
		"CREATEROLESTMT": (*CreateRoleStmtMarshalAlias)(&node),
	})
}

func (node *CreateRoleStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateRoleStmt) Deparse() string {
	panic("Not Implemented")
}
