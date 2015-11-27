// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * "Special join" info.
 *
 * One-sided outer joins constrain the order of joining partially but not
 * completely.  We flatten such joins into the planner's top-level list of
 * relations to join, but record information about each outer join in a
 * SpecialJoinInfo struct.  These structs are kept in the PlannerInfo node's
 * join_info_list.
 *
 * Similarly, semijoins and antijoins created by flattening IN (subselect)
 * and EXISTS(subselect) clauses create partial constraints on join order.
 * These are likewise recorded in SpecialJoinInfo structs.
 *
 * We make SpecialJoinInfos for FULL JOINs even though there is no flexibility
 * of planning for them, because this simplifies make_join_rel()'s API.
 *
 * min_lefthand and min_righthand are the sets of base relids that must be
 * available on each side when performing the special join.  lhs_strict is
 * true if the special join's condition cannot succeed when the LHS variables
 * are all NULL (this means that an outer join can commute with upper-level
 * outer joins even if it appears in their RHS).  We don't bother to set
 * lhs_strict for FULL JOINs, however.
 *
 * It is not valid for either min_lefthand or min_righthand to be empty sets;
 * if they were, this would break the logic that enforces join order.
 *
 * syn_lefthand and syn_righthand are the sets of base relids that are
 * syntactically below this special join.  (These are needed to help compute
 * min_lefthand and min_righthand for higher joins.)
 *
 * delay_upper_joins is set TRUE if we detect a pushed-down clause that has
 * to be evaluated after this join is formed (because it references the RHS).
 * Any outer joins that have such a clause and this join in their RHS cannot
 * commute with this join, because that would leave noplace to check the
 * pushed-down clause.  (We don't track this for FULL JOINs, either.)
 *
 * join_quals is an implicit-AND list of the quals syntactically associated
 * with the join (they may or may not end up being applied at the join level).
 * This is just a side list and does not drive actual application of quals.
 * For JOIN_SEMI joins, this is cleared to NIL in create_unique_path() if
 * the join is found not to be suitable for a uniqueify-the-RHS plan.
 *
 * jointype is never JOIN_RIGHT; a RIGHT JOIN is handled by switching
 * the inputs to make it a LEFT JOIN.  So the allowed values of jointype
 * in a join_info_list member are only LEFT, FULL, SEMI, or ANTI.
 *
 * For purposes of join selectivity estimation, we create transient
 * SpecialJoinInfo structures for regular inner joins; so it is possible
 * to have jointype == JOIN_INNER in such a structure, even though this is
 * not allowed within join_info_list.  We also create transient
 * SpecialJoinInfos with jointype == JOIN_INNER for outer joins, since for
 * cost estimation purposes it is sometimes useful to know the join size under
 * plain innerjoin semantics.  Note that lhs_strict, delay_upper_joins, and
 * join_quals are not set meaningfully within such structs.
 */
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

func (node SpecialJoinInfo) MarshalJSON() ([]byte, error) {
	type SpecialJoinInfoMarshalAlias SpecialJoinInfo
	return json.Marshal(map[string]interface{}{
		"SPECIALJOININFO": (*SpecialJoinInfoMarshalAlias)(&node),
	})
}

func (node *SpecialJoinInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["min_lefthand"] != nil {
		err = json.Unmarshal(fields["min_lefthand"], &node.MinLefthand)
		if err != nil {
			return
		}
	}

	if fields["min_righthand"] != nil {
		err = json.Unmarshal(fields["min_righthand"], &node.MinRighthand)
		if err != nil {
			return
		}
	}

	if fields["syn_lefthand"] != nil {
		err = json.Unmarshal(fields["syn_lefthand"], &node.SynLefthand)
		if err != nil {
			return
		}
	}

	if fields["syn_righthand"] != nil {
		err = json.Unmarshal(fields["syn_righthand"], &node.SynRighthand)
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

	if fields["lhs_strict"] != nil {
		err = json.Unmarshal(fields["lhs_strict"], &node.LhsStrict)
		if err != nil {
			return
		}
	}

	if fields["delay_upper_joins"] != nil {
		err = json.Unmarshal(fields["delay_upper_joins"], &node.DelayUpperJoins)
		if err != nil {
			return
		}
	}

	if fields["join_quals"] != nil {
		node.JoinQuals, err = UnmarshalNodeArrayJSON(fields["join_quals"])
		if err != nil {
			return
		}
	}

	return
}
