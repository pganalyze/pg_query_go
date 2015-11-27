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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["path"] != nil {
		err = json.Unmarshal(fields["path"], &node.Path)
		if err != nil {
			return
		}
	}

	if fields["subpaths"] != nil {
		node.Subpaths, err = UnmarshalNodeArrayJSON(fields["subpaths"])
		if err != nil {
			return
		}
	}

	return
}
