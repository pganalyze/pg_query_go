// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type XmlExpr struct {
	Xpr       Expr          `json:"xpr"`
	Op        XmlExprOp     `json:"op"`         /* xml function ID */
	Name      *string       `json:"name"`       /* name in xml(NAME foo ...) syntaxes */
	NamedArgs []Node        `json:"named_args"` /* non-XML expressions for xml_attributes */
	ArgNames  []Node        `json:"arg_names"`  /* parallel list of Value strings */
	Args      []Node        `json:"args"`       /* list of expressions */
	Xmloption XmlOptionType `json:"xmloption"`  /* DOCUMENT or CONTENT */
	Type      Oid           `json:"type"`       /* target type/typmod for XMLSERIALIZE */
	Typmod    int32         `json:"typmod"`
	Location  int           `json:"location"` /* token location, or -1 if unknown */
}

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
