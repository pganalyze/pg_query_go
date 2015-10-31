// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CteScan) MarshalJSON() ([]byte, error) {
	type CteScanMarshalAlias CteScan
	return json.Marshal(map[string]interface{}{
		"CTESCAN": (*CteScanMarshalAlias)(&node),
	})
}

func (node *CteScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CteScan) Deparse() string {
	panic("Not Implemented")
}
