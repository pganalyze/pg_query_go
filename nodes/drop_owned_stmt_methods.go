// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node DropOwnedStmt) Deparse() string {
	panic("Not Implemented")
}
