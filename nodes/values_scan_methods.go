// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ValuesScan) MarshalJSON() ([]byte, error) {
	type ValuesScanMarshalAlias ValuesScan
	return json.Marshal(map[string]interface{}{
		"VALUESSCAN": (*ValuesScanMarshalAlias)(&node),
	})
}

func (node *ValuesScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ValuesScan) Deparse() string {
	panic("Not Implemented")
}
