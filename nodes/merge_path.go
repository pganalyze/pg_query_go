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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
