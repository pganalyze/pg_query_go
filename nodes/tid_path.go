// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type TidPath struct {
	Path     Path   `json:"path"`
	Tidquals []Node `json:"tidquals"` /* qual(s) involving CTID = something */
}

func (node TidPath) MarshalJSON() ([]byte, error) {
	type TidPathMarshalAlias TidPath
	return json.Marshal(map[string]interface{}{
		"TIDPATH": (*TidPathMarshalAlias)(&node),
	})
}

func (node *TidPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
