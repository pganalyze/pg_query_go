// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateRoleStmt struct {
	StmtType RoleStmtType `json:"stmt_type"` /* ROLE/USER/GROUP */
	Role     *string      `json:"role"`      /* role name */
	Options  []Node       `json:"options"`   /* List of DefElem nodes */
}

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
