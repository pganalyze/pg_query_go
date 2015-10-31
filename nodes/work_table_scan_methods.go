// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node WorkTableScan) MarshalJSON() ([]byte, error) {
	type WorkTableScanMarshalAlias WorkTableScan
	return json.Marshal(map[string]interface{}{
		"WORKTABLESCAN": (*WorkTableScanMarshalAlias)(&node),
	})
}

func (node *WorkTableScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node WorkTableScan) Deparse() string {
	panic("Not Implemented")
}
