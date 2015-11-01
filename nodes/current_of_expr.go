// Auto-generated - DO NOT EDIT

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
