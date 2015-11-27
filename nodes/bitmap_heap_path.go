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

	if fields["bitmapqual"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["bitmapqual"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Path)
			node.Bitmapqual = &val
		}
	}

	return
}
