// Auto-generated - DO NOT EDIT

package pg_query

type IndexScan struct {
	Scan             Scan          `json:"scan"`
	Indexid          Oid           `json:"indexid"`          /* OID of index to scan */
	Indexqual        []Node        `json:"indexqual"`        /* list of index quals (usually OpExprs) */
	Indexqualorig    []Node        `json:"indexqualorig"`    /* the same in original form */
	Indexorderby     []Node        `json:"indexorderby"`     /* list of index ORDER BY exprs */
	Indexorderbyorig []Node        `json:"indexorderbyorig"` /* the same in original form */
	Indexorderdir    ScanDirection `json:"indexorderdir"`    /* forward or backward or don't care */
}
