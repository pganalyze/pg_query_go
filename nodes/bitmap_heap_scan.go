// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type BitmapHeapScan struct {
	Scan           Scan   `json:"scan"`
	Bitmapqualorig []Node `json:"bitmapqualorig"` /* index quals, in standard expr form */
}

func (node BitmapHeapScan) MarshalJSON() ([]byte, error) {
	type BitmapHeapScanMarshalAlias BitmapHeapScan
	return json.Marshal(map[string]interface{}{
		"BITMAPHEAPSCAN": (*BitmapHeapScanMarshalAlias)(&node),
	})
}

func (node *BitmapHeapScan) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["scan"] != nil {
		err = json.Unmarshal(fields["scan"], &node.Scan)
		if err != nil {
			return
		}
	}

	if fields["bitmapqualorig"] != nil {
		node.Bitmapqualorig, err = UnmarshalNodeArrayJSON(fields["bitmapqualorig"])
		if err != nil {
			return
		}
	}

	return
}
