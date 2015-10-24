// Auto-generated - DO NOT EDIT

package pg_query

type AlterRoleStmt struct {
	Role    *string `json:"role"`    /* role name */
	Options []Node  `json:"options"` /* List of DefElem nodes */
	Action  int     `json:"action"`  /* +1 = add members, -1 = drop members */
}
