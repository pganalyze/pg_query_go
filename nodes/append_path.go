// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AppendPath struct {
	Path     Path   `json:"path"`
	Subpaths []Node `json:"subpaths"` /* list of component Paths */
}

func (node AppendPath) MarshalJSON() ([]byte, error) {
	type AppendPathMarshalAlias AppendPath
	return json.Marshal(map[string]interface{}{
		"APPENDPATH": (*AppendPathMarshalAlias)(&node),
	})
}

func (node *AppendPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
