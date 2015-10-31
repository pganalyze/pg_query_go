// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Material) MarshalJSON() ([]byte, error) {
	type MaterialMarshalAlias Material
	return json.Marshal(map[string]interface{}{
		"MATERIAL": (*MaterialMarshalAlias)(&node),
	})
}

func (node *Material) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Material) Deparse() string {
	panic("Not Implemented")
}
