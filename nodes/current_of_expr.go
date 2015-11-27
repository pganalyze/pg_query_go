// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CurrentOfExpr struct {
	Xpr         Expr    `json:"xpr"`
	Cvarno      Index   `json:"cvarno"`       /* RT index of target relation */
	CursorName  *string `json:"cursor_name"`  /* name of referenced cursor, or NULL */
	CursorParam int     `json:"cursor_param"` /* refcursor parameter number, or 0 */
}

func (node CurrentOfExpr) MarshalJSON() ([]byte, error) {
	type CurrentOfExprMarshalAlias CurrentOfExpr
	return json.Marshal(map[string]interface{}{
		"CURRENTOFEXPR": (*CurrentOfExprMarshalAlias)(&node),
	})
}

func (node *CurrentOfExpr) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["cvarno"] != nil {
		err = json.Unmarshal(fields["cvarno"], &node.Cvarno)
		if err != nil {
			return
		}
	}

	if fields["cursor_name"] != nil {
		err = json.Unmarshal(fields["cursor_name"], &node.CursorName)
		if err != nil {
			return
		}
	}

	if fields["cursor_param"] != nil {
		err = json.Unmarshal(fields["cursor_param"], &node.CursorParam)
		if err != nil {
			return
		}
	}

	return
}
