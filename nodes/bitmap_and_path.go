// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["path"] != nil {
		err = json.Unmarshal(fields["path"], &node.Path)
		if err != nil {
			return
		}
	}

	if fields["bitmapquals"] != nil {
		node.Bitmapquals, err = UnmarshalNodeArrayJSON(fields["bitmapquals"])
		if err != nil {
			return
		}
	}

	if fields["bitmapselectivity"] != nil {
		err = json.Unmarshal(fields["bitmapselectivity"], &node.Bitmapselectivity)
		if err != nil {
			return
		}
	}

	return
}
