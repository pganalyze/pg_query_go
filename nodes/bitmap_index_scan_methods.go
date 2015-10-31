// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node BitmapIndexScan) MarshalJSON() ([]byte, error) {
	type BitmapIndexScanMarshalAlias BitmapIndexScan
	return json.Marshal(map[string]interface{}{
		"BITMAPINDEXSCAN": (*BitmapIndexScanMarshalAlias)(&node),
	})
}

func (node *BitmapIndexScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node BitmapIndexScan) Deparse() string {
	panic("Not Implemented")
}
