// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type BitmapAndPath struct {
	Path              Path        `json:"path"`
	Bitmapquals       []Node      `json:"bitmapquals"` /* IndexPaths and BitmapOrPaths */
	Bitmapselectivity Selectivity `json:"bitmapselectivity"`
}

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
