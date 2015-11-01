// Auto-generated - DO NOT EDIT

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
