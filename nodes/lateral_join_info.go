// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
