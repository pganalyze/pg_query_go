// Auto-generated - DO NOT EDIT

package pg_query

type IndexPath struct {
	Path             Path          `json:"path"`
	Indexinfo        *IndexOptInfo `json:"indexinfo"`
	Indexclauses     []Node        `json:"indexclauses"`
	Indexquals       []Node        `json:"indexquals"`
	Indexqualcols    []Node        `json:"indexqualcols"`
	Indexorderbys    []Node        `json:"indexorderbys"`
	Indexorderbycols []Node        `json:"indexorderbycols"`
	Indexscandir     ScanDirection `json:"indexscandir"`
	Indextotalcost   Cost          `json:"indextotalcost"`
	Indexselectivity Selectivity   `json:"indexselectivity"`
}
