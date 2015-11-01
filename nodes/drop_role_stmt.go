// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DropRoleStmt struct {
	Roles     []Node `json:"roles"`      /* List of roles to remove */
	MissingOk bool   `json:"missing_ok"` /* skip error if a role is missing? */
}

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
