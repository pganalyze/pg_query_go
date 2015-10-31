// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node NamedArgExpr) MarshalJSON() ([]byte, error) {
	type NamedArgExprMarshalAlias NamedArgExpr
	return json.Marshal(map[string]interface{}{
		"NAMEDARGEXPR": (*NamedArgExprMarshalAlias)(&node),
	})
}

func (node *NamedArgExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node NamedArgExpr) Deparse() string {
	panic("Not Implemented")
}
