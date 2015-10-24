// Auto-generated - DO NOT EDIT

package pg_query

type AlterDefaultPrivilegesStmt struct {
	Options []Node     `json:"options"` /* list of DefElem */
	Action  *GrantStmt `json:"action"`  /* GRANT/REVOKE action (with objects=NIL) */
}
