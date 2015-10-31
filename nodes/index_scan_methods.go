// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node IndexScan) Deparse() string {
	panic("Not Implemented")
}
