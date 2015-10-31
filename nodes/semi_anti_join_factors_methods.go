// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node SemiAntiJoinFactors) Deparse() string {
	panic("Not Implemented")
}
