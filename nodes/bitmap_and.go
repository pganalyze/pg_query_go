// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type BitmapAnd struct {
	Plan        Plan   `json:"plan"`
	Bitmapplans []Node `json:"bitmapplans"`
}

func (node BitmapAnd) MarshalJSON() ([]byte, error) {
	type BitmapAndMarshalAlias BitmapAnd
	return json.Marshal(map[string]interface{}{
		"BITMAPAND": (*BitmapAndMarshalAlias)(&node),
	})
}

func (node *BitmapAnd) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["bitmapplans"] != nil {
		node.Bitmapplans, err = UnmarshalNodeArrayJSON(fields["bitmapplans"])
		if err != nil {
			return
		}
	}

	return
}
