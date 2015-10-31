// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node BitmapHeapPath) Deparse() string {
	panic("Not Implemented")
}
