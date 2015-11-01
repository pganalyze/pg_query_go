// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterDefaultPrivilegesStmt struct {
	Options []Node     `json:"options"` /* list of DefElem */
	Action  *GrantStmt `json:"action"`  /* GRANT/REVOKE action (with objects=NIL) */
}

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
