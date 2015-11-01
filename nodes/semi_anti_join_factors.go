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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
