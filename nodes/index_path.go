// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node IndexPath) MarshalJSON() ([]byte, error) {
	type IndexPathMarshalAlias IndexPath
	return json.Marshal(map[string]interface{}{
		"INDEXPATH": (*IndexPathMarshalAlias)(&node),
	})
}

func (node *IndexPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
