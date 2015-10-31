// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CollateExpr) MarshalJSON() ([]byte, error) {
	type CollateExprMarshalAlias CollateExpr
	return json.Marshal(map[string]interface{}{
		"COLLATE": (*CollateExprMarshalAlias)(&node),
	})
}

func (node *CollateExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CollateExpr) Deparse() string {
	panic("Not Implemented")
}
