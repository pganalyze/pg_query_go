// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type TidScan struct {
	Scan     Scan   `json:"scan"`
	Tidquals []Node `json:"tidquals"` /* qual(s) involving CTID = something */
}

func (node TidScan) MarshalJSON() ([]byte, error) {
	type TidScanMarshalAlias TidScan
	return json.Marshal(map[string]interface{}{
		"TIDSCAN": (*TidScanMarshalAlias)(&node),
	})
}

func (node *TidScan) UnmarshalJSON(input []byte) (err error) {
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

	if fields["tidquals"] != nil {
		node.Tidquals, err = UnmarshalNodeArrayJSON(fields["tidquals"])
		if err != nil {
			return
		}
	}

	return
}
