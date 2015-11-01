// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type A_Indices struct {
	Lidx Node `json:"lidx"` /* NULL if it's a single subscript */
	Uidx Node `json:"uidx"`
}

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
