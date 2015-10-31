// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Scan) MarshalJSON() ([]byte, error) {
	type ScanMarshalAlias Scan
	return json.Marshal(map[string]interface{}{
		"SCAN": (*ScanMarshalAlias)(&node),
	})
}

func (node *Scan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Scan) Deparse() string {
	panic("Not Implemented")
}
