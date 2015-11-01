// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type UnlistenStmt struct {
	Conditionname *string `json:"conditionname"` /* name to unlisten on, or NULL for all */
}

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
