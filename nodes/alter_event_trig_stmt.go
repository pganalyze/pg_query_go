// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterEventTrigStmt struct {
	Trigname  *string `json:"trigname"`  /* TRIGGER's name */
	Tgenabled byte    `json:"tgenabled"` /* trigger's firing configuration WRT
	 * session_replication_role */
}

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
