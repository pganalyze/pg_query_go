// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Result struct {
	Plan            Plan `json:"plan"`
	Resconstantqual Node `json:"resconstantqual"`
}

func (node Result) MarshalJSON() ([]byte, error) {
	type ResultMarshalAlias Result
	return json.Marshal(map[string]interface{}{
		"RESULT": (*ResultMarshalAlias)(&node),
	})
}

func (node *Result) UnmarshalJSON(input []byte) (err error) {
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

	if fields["resconstantqual"] != nil {
		node.Resconstantqual, err = UnmarshalNodeJSON(fields["resconstantqual"])
		if err != nil {
			return
		}
	}

	return
}
