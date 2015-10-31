// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node TypeName) MarshalJSON() ([]byte, error) {
	type TypeNameMarshalAlias TypeName
	return json.Marshal(map[string]interface{}{
		"TYPENAME": (*TypeNameMarshalAlias)(&node),
	})
}

func (node *TypeName) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node TypeName) Deparse() string {
	panic("Not Implemented")
}
