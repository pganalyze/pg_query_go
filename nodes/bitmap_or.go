// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type BitmapOr struct {
	Plan        Plan   `json:"plan"`
	Bitmapplans []Node `json:"bitmapplans"`
}

func (node BitmapOr) MarshalJSON() ([]byte, error) {
	type BitmapOrMarshalAlias BitmapOr
	return json.Marshal(map[string]interface{}{
		"BITMAPOR": (*BitmapOrMarshalAlias)(&node),
	})
}

func (node *BitmapOr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
