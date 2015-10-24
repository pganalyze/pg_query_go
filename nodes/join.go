// Auto-generated - DO NOT EDIT

package pg_query

type Join struct {
	Plan     Plan     `json:"plan"`
	Jointype JoinType `json:"jointype"`
	Joinqual []Node   `json:"joinqual"` /* JOIN quals (in addition to plan.qual) */
}
