// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node GrantRoleStmt) MarshalJSON() ([]byte, error) {
	type GrantRoleStmtMarshalAlias GrantRoleStmt
	return json.Marshal(map[string]interface{}{
		"GRANTROLESTMT": (*GrantRoleStmtMarshalAlias)(&node),
	})
}

func (node *GrantRoleStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node GrantRoleStmt) Deparse() string {
	panic("Not Implemented")
}
