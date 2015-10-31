// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node CurrentOfExpr) Deparse() string {
	panic("Not Implemented")
}
