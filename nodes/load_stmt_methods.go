// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node LoadStmt) MarshalJSON() ([]byte, error) {
	type LoadStmtMarshalAlias LoadStmt
	return json.Marshal(map[string]interface{}{
		"LOADSTMT": (*LoadStmtMarshalAlias)(&node),
	})
}

func (node *LoadStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node LoadStmt) Deparse() string {
	panic("Not Implemented")
}
