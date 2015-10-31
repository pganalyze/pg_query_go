// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node Limit) Deparse() string {
	panic("Not Implemented")
}
