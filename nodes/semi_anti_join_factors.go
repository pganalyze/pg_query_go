// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * When making cost estimates for a SEMI or ANTI join, there are some
 * correction factors that are needed in both nestloop and hash joins
 * to account for the fact that the executor can stop scanning inner rows
 * as soon as it finds a match to the current outer row.  These numbers
 * depend only on the selected outer and inner join relations, not on the
 * particular paths used for them, so it's worthwhile to calculate them
 * just once per relation pair not once per considered path.  This struct
 * is filled by compute_semi_anti_join_factors and must be passed along
 * to the join cost estimation functions.
 *
 * outer_match_frac is the fraction of the outer tuples that are
 *		expected to have at least one match.
 * match_count is the average number of matches expected for
 *		outer tuples that have at least one match.
 */
type SemiAntiJoinFactors struct {
	OuterMatchFrac Selectivity `json:"outer_match_frac"`
	MatchCount     Selectivity `json:"match_count"`
}

func (node SemiAntiJoinFactors) MarshalJSON() ([]byte, error) {
	type SemiAntiJoinFactorsMarshalAlias SemiAntiJoinFactors
	return json.Marshal(map[string]interface{}{
		"SEMIANTIJOINFACTORS": (*SemiAntiJoinFactorsMarshalAlias)(&node),
	})
}

func (node *SemiAntiJoinFactors) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["outer_match_frac"] != nil {
		err = json.Unmarshal(fields["outer_match_frac"], &node.OuterMatchFrac)
		if err != nil {
			return
		}
	}

	if fields["match_count"] != nil {
		err = json.Unmarshal(fields["match_count"], &node.MatchCount)
		if err != nil {
			return
		}
	}

	return
}
