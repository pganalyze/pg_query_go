// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type FunctionScan struct {
	Scan           Scan   `json:"scan"`
	Functions      []Node `json:"functions"`      /* list of RangeTblFunction nodes */
	Funcordinality bool   `json:"funcordinality"` /* WITH ORDINALITY */
}

func (node FunctionScan) MarshalJSON() ([]byte, error) {
	type FunctionScanMarshalAlias FunctionScan
	return json.Marshal(map[string]interface{}{
		"FUNCTIONSCAN": (*FunctionScanMarshalAlias)(&node),
	})
}

func (node *FunctionScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
