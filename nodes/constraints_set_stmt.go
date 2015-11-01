// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ConstraintsSetStmt struct {
	Constraints []Node `json:"constraints"` /* List of names as RangeVars */
	Deferred    bool   `json:"deferred"`
}

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
