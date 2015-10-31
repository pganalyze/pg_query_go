// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CommonTableExpr) MarshalJSON() ([]byte, error) {
	type CommonTableExprMarshalAlias CommonTableExpr
	return json.Marshal(map[string]interface{}{
		"COMMONTABLEEXPR": (*CommonTableExprMarshalAlias)(&node),
	})
}

func (node *CommonTableExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CommonTableExpr) Deparse() string {
	panic("Not Implemented")
}
