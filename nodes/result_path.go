// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ResultPath struct {
	Path  Path   `json:"path"`
	Quals []Node `json:"quals"`
}

func (node ResultPath) MarshalJSON() ([]byte, error) {
	type ResultPathMarshalAlias ResultPath
	return json.Marshal(map[string]interface{}{
		"RESULTPATH": (*ResultPathMarshalAlias)(&node),
	})
}

func (node *ResultPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
