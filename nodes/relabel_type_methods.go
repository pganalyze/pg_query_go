// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RelabelType) MarshalJSON() ([]byte, error) {
	type RelabelTypeMarshalAlias RelabelType
	return json.Marshal(map[string]interface{}{
		"RELABELTYPE": (*RelabelTypeMarshalAlias)(&node),
	})
}

func (node *RelabelType) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RelabelType) Deparse() string {
	panic("Not Implemented")
}
