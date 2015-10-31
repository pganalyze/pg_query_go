// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CaseExpr) MarshalJSON() ([]byte, error) {
	type CaseExprMarshalAlias CaseExpr
	return json.Marshal(map[string]interface{}{
		"CASE": (*CaseExprMarshalAlias)(&node),
	})
}

func (node *CaseExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CaseExpr) Deparse() string {
	panic("Not Implemented")
}
