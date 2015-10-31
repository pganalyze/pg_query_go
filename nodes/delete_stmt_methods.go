// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DeleteStmt) MarshalJSON() ([]byte, error) {
	type DeleteStmtMarshalAlias DeleteStmt
	return json.Marshal(map[string]interface{}{
		"DELETE FROM": (*DeleteStmtMarshalAlias)(&node),
	})
}

func (node *DeleteStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DeleteStmt) Deparse() string {
	panic("Not Implemented")
}
