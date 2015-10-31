// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SelectStmt) MarshalJSON() ([]byte, error) {
	type SelectStmtMarshalAlias SelectStmt
	return json.Marshal(map[string]interface{}{
		"SELECT": (*SelectStmtMarshalAlias)(&node),
	})
}

func (node *SelectStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SelectStmt) Deparse() string {
	panic("Not Implemented")
}
