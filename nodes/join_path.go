// Auto-generated - DO NOT EDIT

package pg_query

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
