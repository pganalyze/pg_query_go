// Auto-generated - DO NOT EDIT

package pg_query

type CreateEventTrigStmt struct {
	Trigname   *string `json:"trigname"`   /* TRIGGER's name */
	Eventname  *string `json:"eventname"`  /* event's identifier */
	Whenclause []Node  `json:"whenclause"` /* list of DefElems indicating filtering */
	Funcname   []Node  `json:"funcname"`   /* qual. name of function to call */
}
