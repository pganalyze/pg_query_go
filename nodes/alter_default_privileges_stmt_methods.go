// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterDefaultPrivilegesStmt) MarshalJSON() ([]byte, error) {
	type AlterDefaultPrivilegesStmtMarshalAlias AlterDefaultPrivilegesStmt
	return json.Marshal(map[string]interface{}{
		"ALTERDEFAULTPRIVILEGESSTMT": (*AlterDefaultPrivilegesStmtMarshalAlias)(&node),
	})
}

func (node *AlterDefaultPrivilegesStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterDefaultPrivilegesStmt) Deparse() string {
	panic("Not Implemented")
}
