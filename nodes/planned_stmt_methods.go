// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PlannedStmt) MarshalJSON() ([]byte, error) {
	type PlannedStmtMarshalAlias PlannedStmt
	return json.Marshal(map[string]interface{}{
		"PLANNEDSTMT": (*PlannedStmtMarshalAlias)(&node),
	})
}

func (node *PlannedStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PlannedStmt) Deparse() string {
	panic("Not Implemented")
}
