// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *		ValuesScan node
 * ----------------
 */
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

	if fields["values_lists"] != nil {
		node.ValuesLists, err = UnmarshalNodeArrayJSON(fields["values_lists"])
		if err != nil {
			return
		}
	}

	return
}
