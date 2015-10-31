// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node BitmapAndPath) MarshalJSON() ([]byte, error) {
	type BitmapAndPathMarshalAlias BitmapAndPath
	return json.Marshal(map[string]interface{}{
		"BITMAPANDPATH": (*BitmapAndPathMarshalAlias)(&node),
	})
}

func (node *BitmapAndPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node BitmapAndPath) Deparse() string {
	panic("Not Implemented")
}
