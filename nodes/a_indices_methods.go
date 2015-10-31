// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node A_Indices) MarshalJSON() ([]byte, error) {
	type A_IndicesMarshalAlias A_Indices
	return json.Marshal(map[string]interface{}{
		"A_INDICES": (*A_IndicesMarshalAlias)(&node),
	})
}

func (node *A_Indices) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node A_Indices) Deparse() string {
	panic("Not Implemented")
}
