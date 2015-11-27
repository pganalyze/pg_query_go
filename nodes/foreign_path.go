// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ForeignPath struct {
	Path       Path   `json:"path"`
	FdwPrivate []Node `json:"fdw_private"`
}

func (node ForeignPath) MarshalJSON() ([]byte, error) {
	type ForeignPathMarshalAlias ForeignPath
	return json.Marshal(map[string]interface{}{
		"FOREIGNPATH": (*ForeignPathMarshalAlias)(&node),
	})
}

func (node *ForeignPath) UnmarshalJSON(input []byte) (err error) {
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

	if fields["fdw_private"] != nil {
		node.FdwPrivate, err = UnmarshalNodeArrayJSON(fields["fdw_private"])
		if err != nil {
			return
		}
	}

	return
}
