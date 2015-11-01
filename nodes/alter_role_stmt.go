// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterRoleStmt struct {
	Role    *string `json:"role"`    /* role name */
	Options []Node  `json:"options"` /* List of DefElem nodes */
	Action  int     `json:"action"`  /* +1 = add members, -1 = drop members */
}

func (node AlterRoleStmt) MarshalJSON() ([]byte, error) {
	type AlterRoleStmtMarshalAlias AlterRoleStmt
	return json.Marshal(map[string]interface{}{
		"ALTERROLESTMT": (*AlterRoleStmtMarshalAlias)(&node),
	})
}

func (node *AlterRoleStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
