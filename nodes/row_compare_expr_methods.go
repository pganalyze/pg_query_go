// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RowCompareExpr) MarshalJSON() ([]byte, error) {
	type RowCompareExprMarshalAlias RowCompareExpr
	return json.Marshal(map[string]interface{}{
		"ROWCOMPARE": (*RowCompareExprMarshalAlias)(&node),
	})
}

func (node *RowCompareExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RowCompareExpr) Deparse() string {
	panic("Not Implemented")
}
