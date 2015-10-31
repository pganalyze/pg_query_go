// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ConvertRowtypeExpr) MarshalJSON() ([]byte, error) {
	type ConvertRowtypeExprMarshalAlias ConvertRowtypeExpr
	return json.Marshal(map[string]interface{}{
		"CONVERTROWTYPEEXPR": (*ConvertRowtypeExprMarshalAlias)(&node),
	})
}

func (node *ConvertRowtypeExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ConvertRowtypeExpr) Deparse() string {
	panic("Not Implemented")
}
