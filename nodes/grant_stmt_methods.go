// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node GrantStmt) MarshalJSON() ([]byte, error) {
	type GrantStmtMarshalAlias GrantStmt
	return json.Marshal(map[string]interface{}{
		"GRANTSTMT": (*GrantStmtMarshalAlias)(&node),
	})
}

func (node *GrantStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node GrantStmt) Deparse() string {
	panic("Not Implemented")
}
