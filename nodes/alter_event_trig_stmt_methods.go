// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterEventTrigStmt) MarshalJSON() ([]byte, error) {
	type AlterEventTrigStmtMarshalAlias AlterEventTrigStmt
	return json.Marshal(map[string]interface{}{
		"ALTEREVENTTRIGSTMT": (*AlterEventTrigStmtMarshalAlias)(&node),
	})
}

func (node *AlterEventTrigStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterEventTrigStmt) Deparse() string {
	panic("Not Implemented")
}
