// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node ForeignScan) Deparse() string {
	panic("Not Implemented")
}
