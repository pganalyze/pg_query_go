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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	return
}
