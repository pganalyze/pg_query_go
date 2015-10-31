// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node ReassignOwnedStmt) Deparse() string {
	panic("Not Implemented")
}
