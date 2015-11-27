// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["scan"] != nil {
		err = json.Unmarshal(fields["scan"], &node.Scan)
		if err != nil {
			return
		}
	}

	if fields["functions"] != nil {
		node.Functions, err = UnmarshalNodeArrayJSON(fields["functions"])
		if err != nil {
			return
		}
	}

	if fields["funcordinality"] != nil {
		err = json.Unmarshal(fields["funcordinality"], &node.Funcordinality)
		if err != nil {
			return
		}
	}

	return
}
