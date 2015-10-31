// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node MinMaxExpr) MarshalJSON() ([]byte, error) {
	type MinMaxExprMarshalAlias MinMaxExpr
	return json.Marshal(map[string]interface{}{
		"MINMAX": (*MinMaxExprMarshalAlias)(&node),
	})
}

func (node *MinMaxExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node MinMaxExpr) Deparse() string {
	panic("Not Implemented")
}
