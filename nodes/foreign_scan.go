// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ForeignScan struct {
	Scan        Scan   `json:"scan"`
	FdwExprs    []Node `json:"fdw_exprs"`   /* expressions that FDW may evaluate */
	FdwPrivate  []Node `json:"fdw_private"` /* private data for FDW */
	FsSystemCol bool   `json:"fsSystemCol"` /* true if any "system column" is needed */
}

func (node ForeignScan) MarshalJSON() ([]byte, error) {
	type ForeignScanMarshalAlias ForeignScan
	return json.Marshal(map[string]interface{}{
		"FOREIGNSCAN": (*ForeignScanMarshalAlias)(&node),
	})
}

func (node *ForeignScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
