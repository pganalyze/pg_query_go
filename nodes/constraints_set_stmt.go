// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		SET CONSTRAINTS Statement
 * ----------------------
 */
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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["constraints"] != nil {
		node.Constraints, err = UnmarshalNodeArrayJSON(fields["constraints"])
		if err != nil {
			return
		}
	}

	if fields["deferred"] != nil {
		err = json.Unmarshal(fields["deferred"], &node.Deferred)
		if err != nil {
			return
		}
	}

	return
}
