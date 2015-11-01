// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Expr struct {
}

func (node Expr) MarshalJSON() ([]byte, error) {
	type ExprMarshalAlias Expr
	return json.Marshal(map[string]interface{}{
		"EXPR": (*ExprMarshalAlias)(&node),
	})
}

func (node *Expr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
