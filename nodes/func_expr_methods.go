// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node FuncExpr) MarshalJSON() ([]byte, error) {
	type FuncExprMarshalAlias FuncExpr
	return json.Marshal(map[string]interface{}{
		"FUNCEXPR": (*FuncExprMarshalAlias)(&node),
	})
}

func (node *FuncExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node FuncExpr) Deparse() string {
	panic("Not Implemented")
}
