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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["scanrelid"] != nil {
		err = json.Unmarshal(fields["scanrelid"], &node.Scanrelid)
		if err != nil {
			return
		}
	}

	return
}
