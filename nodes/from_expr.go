// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type FromExpr struct {
	Fromlist []Node `json:"fromlist"` /* List of join subtrees */
	Quals    Node   `json:"quals"`    /* qualifiers on join, if any */
}

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
