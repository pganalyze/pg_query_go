// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DropRoleStmt) MarshalJSON() ([]byte, error) {
	type DropRoleStmtMarshalAlias DropRoleStmt
	return json.Marshal(map[string]interface{}{
		"DROPROLESTMT": (*DropRoleStmtMarshalAlias)(&node),
	})
}

func (node *DropRoleStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DropRoleStmt) Deparse() string {
	panic("Not Implemented")
}
