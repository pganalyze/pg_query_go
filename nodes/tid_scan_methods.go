// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node TidScan) MarshalJSON() ([]byte, error) {
	type TidScanMarshalAlias TidScan
	return json.Marshal(map[string]interface{}{
		"TIDSCAN": (*TidScanMarshalAlias)(&node),
	})
}

func (node *TidScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node TidScan) Deparse() string {
	panic("Not Implemented")
}
