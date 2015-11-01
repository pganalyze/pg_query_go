// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type BoolExpr struct {
	Xpr      Expr         `json:"xpr"`
	Boolop   BoolExprType `json:"boolop"`
	Args     []Node       `json:"args"`     /* arguments to this expression */
	Location int          `json:"location"` /* token location, or -1 if unknown */
}

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
