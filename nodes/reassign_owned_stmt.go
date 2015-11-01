// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ReassignOwnedStmt struct {
	Roles   []Node  `json:"roles"`
	Newrole *string `json:"newrole"`
}

func (node ReassignOwnedStmt) MarshalJSON() ([]byte, error) {
	type ReassignOwnedStmtMarshalAlias ReassignOwnedStmt
	return json.Marshal(map[string]interface{}{
		"REASSIGNOWNEDSTMT": (*ReassignOwnedStmtMarshalAlias)(&node),
	})
}

func (node *ReassignOwnedStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
