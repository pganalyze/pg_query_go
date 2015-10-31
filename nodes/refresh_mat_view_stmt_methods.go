// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RefreshMatViewStmt) MarshalJSON() ([]byte, error) {
	type RefreshMatViewStmtMarshalAlias RefreshMatViewStmt
	return json.Marshal(map[string]interface{}{
		"REFRESHMATVIEWSTMT": (*RefreshMatViewStmtMarshalAlias)(&node),
	})
}

func (node *RefreshMatViewStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RefreshMatViewStmt) Deparse() string {
	panic("Not Implemented")
}
