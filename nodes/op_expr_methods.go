// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node OpExpr) MarshalJSON() ([]byte, error) {
	type OpExprMarshalAlias OpExpr
	return json.Marshal(map[string]interface{}{
		"OPEXPR": (*OpExprMarshalAlias)(&node),
	})
}

func (node *OpExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node OpExpr) Deparse() string {
	panic("Not Implemented")
}
