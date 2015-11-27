// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type TidPath struct {
	Path     Path   `json:"path"`
	Tidquals []Node `json:"tidquals"` /* qual(s) involving CTID = something */
}

func (node TidPath) MarshalJSON() ([]byte, error) {
	type TidPathMarshalAlias TidPath
	return json.Marshal(map[string]interface{}{
		"TIDPATH": (*TidPathMarshalAlias)(&node),
	})
}

func (node *TidPath) UnmarshalJSON(input []byte) (err error) {
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

	if fields["tidquals"] != nil {
		node.Tidquals, err = UnmarshalNodeArrayJSON(fields["tidquals"])
		if err != nil {
			return
		}
	}

	return
}
