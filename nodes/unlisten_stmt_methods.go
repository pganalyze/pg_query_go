// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node UnlistenStmt) MarshalJSON() ([]byte, error) {
	type UnlistenStmtMarshalAlias UnlistenStmt
	return json.Marshal(map[string]interface{}{
		"UNLISTENSTMT": (*UnlistenStmtMarshalAlias)(&node),
	})
}

func (node *UnlistenStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node UnlistenStmt) Deparse() string {
	panic("Not Implemented")
}
