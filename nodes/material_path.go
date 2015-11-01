// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type MaterialPath struct {
	Path    Path  `json:"path"`
	Subpath *Path `json:"subpath"`
}

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
