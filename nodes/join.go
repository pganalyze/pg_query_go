// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type Join struct {
	Plan     Plan     `json:"plan"`
	Jointype JoinType `json:"jointype"`
	Joinqual []Node   `json:"joinqual"` /* JOIN quals (in addition to plan.qual) */
}

func (node Join) MarshalJSON() ([]byte, error) {
	type JoinMarshalAlias Join
	return json.Marshal(map[string]interface{}{
		"JOIN": (*JoinMarshalAlias)(&node),
	})
}

func (node *Join) UnmarshalJSON(input []byte) (err error) {
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

	if fields["jointype"] != nil {
		err = json.Unmarshal(fields["jointype"], &node.Jointype)
		if err != nil {
			return
		}
	}

	if fields["joinqual"] != nil {
		node.Joinqual, err = UnmarshalNodeArrayJSON(fields["joinqual"])
		if err != nil {
			return
		}
	}

	return
}
