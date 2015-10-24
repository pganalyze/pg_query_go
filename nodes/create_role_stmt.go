// Auto-generated - DO NOT EDIT

package pg_query

type CreateRoleStmt struct {
	StmtType RoleStmtType `json:"stmt_type"` /* ROLE/USER/GROUP */
	Role     *string      `json:"role"`      /* role name */
	Options  []Node       `json:"options"`   /* List of DefElem nodes */
}
