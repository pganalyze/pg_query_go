// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node VariableSetStmt) MarshalJSON() ([]byte, error) {
	type VariableSetStmtMarshalAlias VariableSetStmt
	return json.Marshal(map[string]interface{}{
		"SET": (*VariableSetStmtMarshalAlias)(&node),
	})
}

func (node *VariableSetStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node VariableSetStmt) Deparse() string {
	panic("Not Implemented")
}
