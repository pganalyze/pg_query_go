// Auto-generated - DO NOT EDIT

package pg_query

type SpecialJoinInfo struct {
	MinLefthand     []uint32 `json:"min_lefthand"`      /* base relids in minimum LHS for join */
	MinRighthand    []uint32 `json:"min_righthand"`     /* base relids in minimum RHS for join */
	SynLefthand     []uint32 `json:"syn_lefthand"`      /* base relids syntactically within LHS */
	SynRighthand    []uint32 `json:"syn_righthand"`     /* base relids syntactically within RHS */
	Jointype        JoinType `json:"jointype"`          /* always INNER, LEFT, FULL, SEMI, or ANTI */
	LhsStrict       bool     `json:"lhs_strict"`        /* joinclause is strict for some LHS rel */
	DelayUpperJoins bool     `json:"delay_upper_joins"` /* can't commute with upper RHS */
	JoinQuals       []Node   `json:"join_quals"`        /* join quals, in implicit-AND list format */
}
