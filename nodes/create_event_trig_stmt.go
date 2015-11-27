// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		Create EVENT TRIGGER Statement
 * ----------------------
 */
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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["trigname"] != nil {
		err = json.Unmarshal(fields["trigname"], &node.Trigname)
		if err != nil {
			return
		}
	}

	if fields["eventname"] != nil {
		err = json.Unmarshal(fields["eventname"], &node.Eventname)
		if err != nil {
			return
		}
	}

	if fields["whenclause"] != nil {
		node.Whenclause, err = UnmarshalNodeArrayJSON(fields["whenclause"])
		if err != nil {
			return
		}
	}

	if fields["funcname"] != nil {
		node.Funcname, err = UnmarshalNodeArrayJSON(fields["funcname"])
		if err != nil {
			return
		}
	}

	return
}
