// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AccessPriv struct {
	PrivName *string `json:"priv_name"` /* string name of privilege */
	Cols     []Node  `json:"cols"`      /* list of Value strings */
}

func (node AccessPriv) MarshalJSON() ([]byte, error) {
	type AccessPrivMarshalAlias AccessPriv
	return json.Marshal(map[string]interface{}{
		"ACCESSPRIV": (*AccessPrivMarshalAlias)(&node),
	})
}

func (node *AccessPriv) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["priv_name"] != nil {
		err = json.Unmarshal(fields["priv_name"], &node.PrivName)
		if err != nil {
			return
		}
	}

	if fields["cols"] != nil {
		node.Cols, err = UnmarshalNodeArrayJSON(fields["cols"])
		if err != nil {
			return
		}
	}

	return
}
