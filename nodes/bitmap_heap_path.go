// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type BitmapHeapPath struct {
	Path       Path  `json:"path"`
	Bitmapqual *Path `json:"bitmapqual"` /* IndexPath, BitmapAndPath, BitmapOrPath */
}

func (node BitmapHeapPath) MarshalJSON() ([]byte, error) {
	type BitmapHeapPathMarshalAlias BitmapHeapPath
	return json.Marshal(map[string]interface{}{
		"BITMAPHEAPPATH": (*BitmapHeapPathMarshalAlias)(&node),
	})
}

func (node *BitmapHeapPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
