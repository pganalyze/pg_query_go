// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Scan struct {
	Plan      Plan  `json:"plan"`
	Scanrelid Index `json:"scanrelid"` /* relid is index into the range table */
}

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
