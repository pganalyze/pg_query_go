// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type VariableSetStmt struct {
	Kind    VariableSetKind `json:"kind"`
	Name    *string         `json:"name"`     /* variable to be set */
	Args    []Node          `json:"args"`     /* List of A_Const nodes */
	IsLocal bool            `json:"is_local"` /* SET LOCAL? */
}

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
