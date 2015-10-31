// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node BoolExpr) MarshalJSON() ([]byte, error) {
	type BoolExprMarshalAlias BoolExpr
	return json.Marshal(map[string]interface{}{
		"BOOLEXPR": (*BoolExprMarshalAlias)(&node),
	})
}

func (node *BoolExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node BoolExpr) Deparse() string {
	panic("Not Implemented")
}
