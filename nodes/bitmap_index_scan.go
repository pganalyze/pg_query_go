// Auto-generated - DO NOT EDIT

package pg_query

type BitmapIndexScan struct {
	Scan          Scan   `json:"scan"`
	Indexid       Oid    `json:"indexid"`       /* OID of index to scan */
	Indexqual     []Node `json:"indexqual"`     /* list of index quals (OpExprs) */
	Indexqualorig []Node `json:"indexqualorig"` /* the same in original form */
}
