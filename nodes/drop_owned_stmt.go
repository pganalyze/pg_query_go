// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DropOwnedStmt struct {
	Roles    []Node       `json:"roles"`
	Behavior DropBehavior `json:"behavior"`
}

func (node DropOwnedStmt) MarshalJSON() ([]byte, error) {
	type DropOwnedStmtMarshalAlias DropOwnedStmt
	return json.Marshal(map[string]interface{}{
		"DROPOWNEDSTMT": (*DropOwnedStmtMarshalAlias)(&node),
	})
}

func (node *DropOwnedStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
