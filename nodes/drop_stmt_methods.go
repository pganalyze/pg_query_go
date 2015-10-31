// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DropStmt) MarshalJSON() ([]byte, error) {
	type DropStmtMarshalAlias DropStmt
	return json.Marshal(map[string]interface{}{
		"DROP": (*DropStmtMarshalAlias)(&node),
	})
}

func (node *DropStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DropStmt) Deparse() string {
	panic("Not Implemented")
}
