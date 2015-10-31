// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node IndexOptInfo) MarshalJSON() ([]byte, error) {
	type IndexOptInfoMarshalAlias IndexOptInfo
	return json.Marshal(map[string]interface{}{
		"INDEXOPTINFO": (*IndexOptInfoMarshalAlias)(&node),
	})
}

func (node *IndexOptInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node IndexOptInfo) Deparse() string {
	panic("Not Implemented")
}
