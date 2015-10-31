// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node WithCheckOption) MarshalJSON() ([]byte, error) {
	type WithCheckOptionMarshalAlias WithCheckOption
	return json.Marshal(map[string]interface{}{
		"WITHCHECKOPTION": (*WithCheckOptionMarshalAlias)(&node),
	})
}

func (node *WithCheckOption) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node WithCheckOption) Deparse() string {
	panic("Not Implemented")
}
