// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * BitmapOrPath represents a BitmapOr plan node; it can only appear as
 * part of the substructure of a BitmapHeapPath.  The Path structure is
 * a bit more heavyweight than we really need for this, but for simplicity
 * we make it a derivative of Path anyway.
 */
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
