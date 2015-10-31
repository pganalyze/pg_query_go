// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RowExpr) MarshalJSON() ([]byte, error) {
	type RowExprMarshalAlias RowExpr
	return json.Marshal(map[string]interface{}{
		"ROW": (*RowExprMarshalAlias)(&node),
	})
}

func (node *RowExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RowExpr) Deparse() string {
	panic("Not Implemented")
}
