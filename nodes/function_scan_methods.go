// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node FunctionScan) Deparse() string {
	panic("Not Implemented")
}
