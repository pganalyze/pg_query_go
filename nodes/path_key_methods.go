// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PathKey) MarshalJSON() ([]byte, error) {
	type PathKeyMarshalAlias PathKey
	return json.Marshal(map[string]interface{}{
		"PATHKEY": (*PathKeyMarshalAlias)(&node),
	})
}

func (node *PathKey) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PathKey) Deparse() string {
	panic("Not Implemented")
}
