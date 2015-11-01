// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Limit struct {
	Plan        Plan `json:"plan"`
	LimitOffset Node `json:"limitOffset"` /* OFFSET parameter, or NULL if none */
	LimitCount  Node `json:"limitCount"`  /* COUNT parameter, or NULL if none */
}

func (node Limit) MarshalJSON() ([]byte, error) {
	type LimitMarshalAlias Limit
	return json.Marshal(map[string]interface{}{
		"LIMIT": (*LimitMarshalAlias)(&node),
	})
}

func (node *Limit) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
