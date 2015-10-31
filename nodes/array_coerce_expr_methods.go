// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ArrayCoerceExpr) MarshalJSON() ([]byte, error) {
	type ArrayCoerceExprMarshalAlias ArrayCoerceExpr
	return json.Marshal(map[string]interface{}{
		"ARRAYCOERCEEXPR": (*ArrayCoerceExprMarshalAlias)(&node),
	})
}

func (node *ArrayCoerceExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ArrayCoerceExpr) Deparse() string {
	panic("Not Implemented")
}
