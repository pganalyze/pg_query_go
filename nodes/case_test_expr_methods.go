// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CaseTestExpr) MarshalJSON() ([]byte, error) {
	type CaseTestExprMarshalAlias CaseTestExpr
	return json.Marshal(map[string]interface{}{
		"CASETESTEXPR": (*CaseTestExprMarshalAlias)(&node),
	})
}

func (node *CaseTestExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CaseTestExpr) Deparse() string {
	panic("Not Implemented")
}
