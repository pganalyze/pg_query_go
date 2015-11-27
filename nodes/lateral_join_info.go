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
