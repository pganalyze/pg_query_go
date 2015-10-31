// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ScalarArrayOpExpr) MarshalJSON() ([]byte, error) {
	type ScalarArrayOpExprMarshalAlias ScalarArrayOpExpr
	return json.Marshal(map[string]interface{}{
		"SCALARARRAYOPEXPR": (*ScalarArrayOpExprMarshalAlias)(&node),
	})
}

func (node *ScalarArrayOpExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ScalarArrayOpExpr) Deparse() string {
	panic("Not Implemented")
}
