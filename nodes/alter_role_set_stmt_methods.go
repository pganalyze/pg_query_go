// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterRoleSetStmt) MarshalJSON() ([]byte, error) {
	type AlterRoleSetStmtMarshalAlias AlterRoleSetStmt
	return json.Marshal(map[string]interface{}{
		"ALTERROLESETSTMT": (*AlterRoleSetStmtMarshalAlias)(&node),
	})
}

func (node *AlterRoleSetStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterRoleSetStmt) Deparse() string {
	panic("Not Implemented")
}
