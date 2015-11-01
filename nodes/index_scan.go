// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type IndexScan struct {
	Scan             Scan          `json:"scan"`
	Indexid          Oid           `json:"indexid"`          /* OID of index to scan */
	Indexqual        []Node        `json:"indexqual"`        /* list of index quals (usually OpExprs) */
	Indexqualorig    []Node        `json:"indexqualorig"`    /* the same in original form */
	Indexorderby     []Node        `json:"indexorderby"`     /* list of index ORDER BY exprs */
	Indexorderbyorig []Node        `json:"indexorderbyorig"` /* the same in original form */
	Indexorderdir    ScanDirection `json:"indexorderdir"`    /* forward or backward or don't care */
}

func (node IndexScan) MarshalJSON() ([]byte, error) {
	type IndexScanMarshalAlias IndexScan
	return json.Marshal(map[string]interface{}{
		"INDEXSCAN": (*IndexScanMarshalAlias)(&node),
	})
}

func (node *IndexScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
