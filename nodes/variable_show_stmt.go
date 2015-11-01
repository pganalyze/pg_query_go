// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type VariableShowStmt struct {
	Name *string `json:"name"`
}

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
