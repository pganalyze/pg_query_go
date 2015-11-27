// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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
