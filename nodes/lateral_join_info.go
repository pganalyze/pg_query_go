// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * "Lateral join" info.
 *
 * Lateral references constrain the join order in a way that's somewhat like
 * outer joins, though different in detail.  We construct a LateralJoinInfo
 * for each lateral cross-reference, placing them in the PlannerInfo node's
 * lateral_info_list.
 *
 * For unflattened LATERAL RTEs, we generate LateralJoinInfo(s) in which
 * lateral_rhs is the relid of the LATERAL baserel, and lateral_lhs is a set
 * of relids of baserels it references, all of which must be present on the
 * LHS to compute a parameter needed by the RHS.  Typically, lateral_lhs is
 * a singleton, but it can include multiple rels if the RHS references a
 * PlaceHolderVar with a multi-rel ph_eval_at level.  We disallow joining to
 * only part of the LHS in such cases, since that would result in a join tree
 * with no convenient place to compute the PHV.
 *
 * When an appendrel contains lateral references (eg "LATERAL (SELECT x.col1
 * UNION ALL SELECT y.col2)"), the LateralJoinInfos reference the parent
 * baserel not the member otherrels, since it is the parent relid that is
 * considered for joining purposes.
 *
 * If any LATERAL RTEs were flattened into the parent query, it is possible
 * that the query now contains PlaceHolderVars containing lateral references,
 * representing expressions that need to be evaluated at particular spots in
 * the jointree but contain lateral references to Vars from elsewhere.  These
 * give rise to LateralJoinInfos in which lateral_rhs is the evaluation point
 * of a PlaceHolderVar and lateral_lhs is the set of lateral rels it needs.
 */
type LateralJoinInfo struct {
	LateralLhs []uint32 `json:"lateral_lhs"` /* rels needed to compute a lateral value */
	LateralRhs []uint32 `json:"lateral_rhs"` /* rel where lateral value is needed */
}

func (node LateralJoinInfo) MarshalJSON() ([]byte, error) {
	type LateralJoinInfoMarshalAlias LateralJoinInfo
	return json.Marshal(map[string]interface{}{
		"LATERALJOININFO": (*LateralJoinInfoMarshalAlias)(&node),
	})
}

func (node *LateralJoinInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["lateral_lhs"] != nil {
		err = json.Unmarshal(fields["lateral_lhs"], &node.LateralLhs)
		if err != nil {
			return
		}
	}

	if fields["lateral_rhs"] != nil {
		err = json.Unmarshal(fields["lateral_rhs"], &node.LateralRhs)
		if err != nil {
			return
		}
	}

	return
}
