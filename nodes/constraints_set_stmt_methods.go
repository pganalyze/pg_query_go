// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ConstraintsSetStmt) MarshalJSON() ([]byte, error) {
	type ConstraintsSetStmtMarshalAlias ConstraintsSetStmt
	return json.Marshal(map[string]interface{}{
		"CONSTRAINTSSETSTMT": (*ConstraintsSetStmtMarshalAlias)(&node),
	})
}

func (node *ConstraintsSetStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ConstraintsSetStmt) Deparse() string {
	panic("Not Implemented")
}
