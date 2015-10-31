// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node FetchStmt) MarshalJSON() ([]byte, error) {
	type FetchStmtMarshalAlias FetchStmt
	return json.Marshal(map[string]interface{}{
		"FETCHSTMT": (*FetchStmtMarshalAlias)(&node),
	})
}

func (node *FetchStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node FetchStmt) Deparse() string {
	panic("Not Implemented")
}
