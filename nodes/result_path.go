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

	if fields["quals"] != nil {
		node.Quals, err = UnmarshalNodeArrayJSON(fields["quals"])
		if err != nil {
			return
		}
	}

	return
}
