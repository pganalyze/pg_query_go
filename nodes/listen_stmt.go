// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ListenStmt struct {
	Conditionname *string `json:"conditionname"` /* condition name to listen on */
}

func (node ListenStmt) MarshalJSON() ([]byte, error) {
	type ListenStmtMarshalAlias ListenStmt
	return json.Marshal(map[string]interface{}{
		"LISTENSTMT": (*ListenStmtMarshalAlias)(&node),
	})
}

func (node *ListenStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
