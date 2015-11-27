// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type Alias struct {
	Aliasname *string `json:"aliasname"` /* aliased rel name (never qualified) */
	Colnames  []Node  `json:"colnames"`  /* optional list of column aliases */
}

func (node Alias) MarshalJSON() ([]byte, error) {
	type AliasMarshalAlias Alias
	return json.Marshal(map[string]interface{}{
		"ALIAS": (*AliasMarshalAlias)(&node),
	})
}

func (node *Alias) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["aliasname"] != nil {
		err = json.Unmarshal(fields["aliasname"], &node.Aliasname)
		if err != nil {
			return
		}
	}

	if fields["colnames"] != nil {
		node.Colnames, err = UnmarshalNodeArrayJSON(fields["colnames"])
		if err != nil {
			return
		}
	}

	return
}
