// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node BitmapOrPath) MarshalJSON() ([]byte, error) {
	type BitmapOrPathMarshalAlias BitmapOrPath
	return json.Marshal(map[string]interface{}{
		"BITMAPORPATH": (*BitmapOrPathMarshalAlias)(&node),
	})
}

func (node *BitmapOrPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node BitmapOrPath) Deparse() string {
	panic("Not Implemented")
}
