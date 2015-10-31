// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node UpdateStmt) MarshalJSON() ([]byte, error) {
	type UpdateStmtMarshalAlias UpdateStmt
	return json.Marshal(map[string]interface{}{
		"UPDATE": (*UpdateStmtMarshalAlias)(&node),
	})
}

func (node *UpdateStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node UpdateStmt) Deparse() string {
	panic("Not Implemented")
}
