// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node BitmapHeapScan) Deparse() string {
	panic("Not Implemented")
}
