// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ValuesScan struct {
	Scan        Scan   `json:"scan"`
	ValuesLists []Node `json:"values_lists"` /* list of expression lists */
}

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
