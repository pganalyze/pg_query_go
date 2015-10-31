// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CoalesceExpr) MarshalJSON() ([]byte, error) {
	type CoalesceExprMarshalAlias CoalesceExpr
	return json.Marshal(map[string]interface{}{
		"COALESCE": (*CoalesceExprMarshalAlias)(&node),
	})
}

func (node *CoalesceExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CoalesceExpr) Deparse() string {
	panic("Not Implemented")
}
