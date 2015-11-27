// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type PrivGrantee struct {
	Rolname *string `json:"rolname"` /* if NULL then PUBLIC */
}

func (node PrivGrantee) MarshalJSON() ([]byte, error) {
	type PrivGranteeMarshalAlias PrivGrantee
	return json.Marshal(map[string]interface{}{
		"PRIVGRANTEE": (*PrivGranteeMarshalAlias)(&node),
	})
}

func (node *PrivGrantee) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["rolname"] != nil {
		err = json.Unmarshal(fields["rolname"], &node.Rolname)
		if err != nil {
			return
		}
	}

	return
}
