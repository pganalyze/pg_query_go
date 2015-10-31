// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node XmlSerialize) Deparse() string {
	panic("Not Implemented")
}
