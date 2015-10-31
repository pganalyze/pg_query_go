// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node LateralJoinInfo) Deparse() string {
	panic("Not Implemented")
}
