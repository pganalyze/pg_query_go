// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SpecialJoinInfo) MarshalJSON() ([]byte, error) {
	type SpecialJoinInfoMarshalAlias SpecialJoinInfo
	return json.Marshal(map[string]interface{}{
		"SPECIALJOININFO": (*SpecialJoinInfoMarshalAlias)(&node),
	})
}

func (node *SpecialJoinInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SpecialJoinInfo) Deparse() string {
	panic("Not Implemented")
}
