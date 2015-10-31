// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Join) MarshalJSON() ([]byte, error) {
	type JoinMarshalAlias Join
	return json.Marshal(map[string]interface{}{
		"JOIN": (*JoinMarshalAlias)(&node),
	})
}

func (node *Join) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Join) Deparse() string {
	panic("Not Implemented")
}
