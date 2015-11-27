// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type RangeTblRef struct {
	Rtindex int `json:"rtindex"`
}

func (node RangeTblRef) MarshalJSON() ([]byte, error) {
	type RangeTblRefMarshalAlias RangeTblRef
	return json.Marshal(map[string]interface{}{
		"RANGETBLREF": (*RangeTblRefMarshalAlias)(&node),
	})
}

func (node *RangeTblRef) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["rtindex"] != nil {
		err = json.Unmarshal(fields["rtindex"], &node.Rtindex)
		if err != nil {
			return
		}
	}

	return
}
