// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type XmlSerialize struct {
	Xmloption XmlOptionType `json:"xmloption"` /* DOCUMENT or CONTENT */
	Expr      Node          `json:"expr"`
	TypeName  *TypeName     `json:"typeName"`
	Location  int           `json:"location"` /* token location, or -1 if unknown */
}

func (node XmlSerialize) MarshalJSON() ([]byte, error) {
	type XmlSerializeMarshalAlias XmlSerialize
	return json.Marshal(map[string]interface{}{
		"XMLSERIALIZE": (*XmlSerializeMarshalAlias)(&node),
	})
}

func (node *XmlSerialize) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
