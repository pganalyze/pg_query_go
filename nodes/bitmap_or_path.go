// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type BitmapOrPath struct {
	Path              Path        `json:"path"`
	Bitmapquals       []Node      `json:"bitmapquals"` /* IndexPaths and BitmapAndPaths */
	Bitmapselectivity Selectivity `json:"bitmapselectivity"`
}

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
