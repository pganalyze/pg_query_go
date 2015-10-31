// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node ListenStmt) Deparse() string {
	panic("Not Implemented")
}
