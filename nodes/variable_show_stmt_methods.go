// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node VariableShowStmt) MarshalJSON() ([]byte, error) {
	type VariableShowStmtMarshalAlias VariableShowStmt
	return json.Marshal(map[string]interface{}{
		"SHOW": (*VariableShowStmtMarshalAlias)(&node),
	})
}

func (node *VariableShowStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node VariableShowStmt) Deparse() string {
	panic("Not Implemented")
}
