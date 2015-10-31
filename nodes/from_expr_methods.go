// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node FromExpr) MarshalJSON() ([]byte, error) {
	type FromExprMarshalAlias FromExpr
	return json.Marshal(map[string]interface{}{
		"FROMEXPR": (*FromExprMarshalAlias)(&node),
	})
}

func (node *FromExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node FromExpr) Deparse() string {
	panic("Not Implemented")
}
