// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterRoleStmt) MarshalJSON() ([]byte, error) {
	type AlterRoleStmtMarshalAlias AlterRoleStmt
	return json.Marshal(map[string]interface{}{
		"ALTERROLESTMT": (*AlterRoleStmtMarshalAlias)(&node),
	})
}

func (node *AlterRoleStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterRoleStmt) Deparse() string {
	panic("Not Implemented")
}
