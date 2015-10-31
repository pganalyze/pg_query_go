// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ArrayExpr) MarshalJSON() ([]byte, error) {
	type ArrayExprMarshalAlias ArrayExpr
	return json.Marshal(map[string]interface{}{
		"ARRAY": (*ArrayExprMarshalAlias)(&node),
	})
}

func (node *ArrayExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ArrayExpr) Deparse() string {
	panic("Not Implemented")
}
