// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type JoinPath struct {
	Path Path `json:"path"`

	Jointype JoinType `json:"jointype"`

	Outerjoinpath *Path `json:"outerjoinpath"` /* path for the outer side of the join */
	Innerjoinpath *Path `json:"innerjoinpath"` /* path for the inner side of the join */

	Joinrestrictinfo []Node `json:"joinrestrictinfo"` /* RestrictInfos to apply to join */

	/*
	 * See the notes for RelOptInfo and ParamPathInfo to understand why
	 * joinrestrictinfo is needed in JoinPath, and can't be merged into the
	 * parent RelOptInfo.
	 */
}

func (node JoinPath) MarshalJSON() ([]byte, error) {
	type JoinPathMarshalAlias JoinPath
	return json.Marshal(map[string]interface{}{
		"JOINPATH": (*JoinPathMarshalAlias)(&node),
	})
}

func (node *JoinPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
