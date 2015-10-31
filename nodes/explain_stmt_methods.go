// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ExplainStmt) MarshalJSON() ([]byte, error) {
	type ExplainStmtMarshalAlias ExplainStmt
	return json.Marshal(map[string]interface{}{
		"EXPLAIN": (*ExplainStmtMarshalAlias)(&node),
	})
}

func (node *ExplainStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ExplainStmt) Deparse() string {
	panic("Not Implemented")
}
