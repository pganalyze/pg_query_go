// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type MergePath struct {
	Jpath            JoinPath `json:"jpath"`
	PathMergeclauses []Node   `json:"path_mergeclauses"` /* join clauses to be used for merge */
	Outersortkeys    []Node   `json:"outersortkeys"`     /* keys for explicit sort, if any */
	Innersortkeys    []Node   `json:"innersortkeys"`     /* keys for explicit sort, if any */
	MaterializeInner bool     `json:"materialize_inner"` /* add Materialize to inner? */
}

func (node MergePath) MarshalJSON() ([]byte, error) {
	type MergePathMarshalAlias MergePath
	return json.Marshal(map[string]interface{}{
		"MERGEPATH": (*MergePathMarshalAlias)(&node),
	})
}

func (node *MergePath) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["jpath"] != nil {
		err = json.Unmarshal(fields["jpath"], &node.Jpath)
		if err != nil {
			return
		}
	}

	if fields["path_mergeclauses"] != nil {
		node.PathMergeclauses, err = UnmarshalNodeArrayJSON(fields["path_mergeclauses"])
		if err != nil {
			return
		}
	}

	if fields["outersortkeys"] != nil {
		node.Outersortkeys, err = UnmarshalNodeArrayJSON(fields["outersortkeys"])
		if err != nil {
			return
		}
	}

	if fields["innersortkeys"] != nil {
		node.Innersortkeys, err = UnmarshalNodeArrayJSON(fields["innersortkeys"])
		if err != nil {
			return
		}
	}

	if fields["materialize_inner"] != nil {
		err = json.Unmarshal(fields["materialize_inner"], &node.MaterializeInner)
		if err != nil {
			return
		}
	}

	return
}
