// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["fromlist"] != nil {
		node.Fromlist, err = UnmarshalNodeArrayJSON(fields["fromlist"])
		if err != nil {
			return
		}
	}

	if fields["quals"] != nil {
		node.Quals, err = UnmarshalNodeJSON(fields["quals"])
		if err != nil {
			return
		}
	}

	return
}
