// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node InsertStmt) MarshalJSON() ([]byte, error) {
	type InsertStmtMarshalAlias InsertStmt
	return json.Marshal(map[string]interface{}{
		"INSERT INTO": (*InsertStmtMarshalAlias)(&node),
	})
}

func (node *InsertStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node InsertStmt) Deparse() string {
	panic("Not Implemented")
}
