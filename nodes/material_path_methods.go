// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node MaterialPath) MarshalJSON() ([]byte, error) {
	type MaterialPathMarshalAlias MaterialPath
	return json.Marshal(map[string]interface{}{
		"MATERIALPATH": (*MaterialPathMarshalAlias)(&node),
	})
}

func (node *MaterialPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node MaterialPath) Deparse() string {
	panic("Not Implemented")
}
