// Auto-generated - DO NOT EDIT

package pg_query

type DropRoleStmt struct {
	Roles     []Node `json:"roles"`      /* List of roles to remove */
	MissingOk bool   `json:"missing_ok"` /* skip error if a role is missing? */
}
