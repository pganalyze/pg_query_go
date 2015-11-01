// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type BitmapIndexScan struct {
	Scan          Scan   `json:"scan"`
	Indexid       Oid    `json:"indexid"`       /* OID of index to scan */
	Indexqual     []Node `json:"indexqual"`     /* list of index quals (OpExprs) */
	Indexqualorig []Node `json:"indexqualorig"` /* the same in original form */
}

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
