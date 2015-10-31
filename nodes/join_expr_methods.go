// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node JoinExpr) MarshalJSON() ([]byte, error) {
	type JoinExprMarshalAlias JoinExpr
	return json.Marshal(map[string]interface{}{
		"JOINEXPR": (*JoinExprMarshalAlias)(&node),
	})
}

func (node *JoinExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node JoinExpr) Deparse() string {
	panic("Not Implemented")
}
