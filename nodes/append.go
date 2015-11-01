// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Append struct {
	Plan        Plan   `json:"plan"`
	Appendplans []Node `json:"appendplans"`
}

func (node Append) MarshalJSON() ([]byte, error) {
	type AppendMarshalAlias Append
	return json.Marshal(map[string]interface{}{
		"APPEND": (*AppendMarshalAlias)(&node),
	})
}

func (node *Append) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
