// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateEventTrigStmt struct {
	Trigname   *string `json:"trigname"`   /* TRIGGER's name */
	Eventname  *string `json:"eventname"`  /* event's identifier */
	Whenclause []Node  `json:"whenclause"` /* list of DefElems indicating filtering */
	Funcname   []Node  `json:"funcname"`   /* qual. name of function to call */
}

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
