// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node XmlExpr) MarshalJSON() ([]byte, error) {
	type XmlExprMarshalAlias XmlExpr
	return json.Marshal(map[string]interface{}{
		"XMLEXPR": (*XmlExprMarshalAlias)(&node),
	})
}

func (node *XmlExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node XmlExpr) Deparse() string {
	panic("Not Implemented")
}
