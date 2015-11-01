// Auto-generated - DO NOT EDIT

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
