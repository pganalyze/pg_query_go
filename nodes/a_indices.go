// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * A_Indices - array subscript or slice bounds ([lidx:uidx] or [uidx])
 */
type A_Indices struct {
	Lidx Node `json:"lidx"` /* NULL if it's a single subscript */
	Uidx Node `json:"uidx"`
}

func (node A_Indices) MarshalJSON() ([]byte, error) {
	type A_IndicesMarshalAlias A_Indices
	return json.Marshal(map[string]interface{}{
		"A_Indices": (*A_IndicesMarshalAlias)(&node),
	})
}

func (node *A_Indices) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["lidx"] != nil {
		node.Lidx, err = UnmarshalNodeJSON(fields["lidx"])
		if err != nil {
			return
		}
	}

	if fields["uidx"] != nil {
		node.Uidx, err = UnmarshalNodeJSON(fields["uidx"])
		if err != nil {
			return
		}
	}

	return
}
