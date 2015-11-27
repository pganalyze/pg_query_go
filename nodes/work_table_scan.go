// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type WorkTableScan struct {
	Scan    Scan `json:"scan"`
	WtParam int  `json:"wtParam"` /* ID of Param representing work table */
}

func (node WorkTableScan) MarshalJSON() ([]byte, error) {
	type WorkTableScanMarshalAlias WorkTableScan
	return json.Marshal(map[string]interface{}{
		"WORKTABLESCAN": (*WorkTableScanMarshalAlias)(&node),
	})
}

func (node *WorkTableScan) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["scan"] != nil {
		err = json.Unmarshal(fields["scan"], &node.Scan)
		if err != nil {
			return
		}
	}

	if fields["wtParam"] != nil {
		err = json.Unmarshal(fields["wtParam"], &node.WtParam)
		if err != nil {
			return
		}
	}

	return
}
