// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RefreshMatViewStmt struct {
	Concurrent bool      `json:"concurrent"` /* allow concurrent access? */
	SkipData   bool      `json:"skipData"`   /* true for WITH NO DATA */
	Relation   *RangeVar `json:"relation"`   /* relation to insert into */
}

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
