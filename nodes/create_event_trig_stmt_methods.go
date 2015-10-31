// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateEventTrigStmt) MarshalJSON() ([]byte, error) {
	type CreateEventTrigStmtMarshalAlias CreateEventTrigStmt
	return json.Marshal(map[string]interface{}{
		"CREATEEVENTTRIGSTMT": (*CreateEventTrigStmtMarshalAlias)(&node),
	})
}

func (node *CreateEventTrigStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateEventTrigStmt) Deparse() string {
	panic("Not Implemented")
}
