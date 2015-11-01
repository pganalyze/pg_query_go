// Auto-generated - DO NOT EDIT

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
