// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AccessPriv) MarshalJSON() ([]byte, error) {
	type AccessPrivMarshalAlias AccessPriv
	return json.Marshal(map[string]interface{}{
		"ACCESSPRIV": (*AccessPrivMarshalAlias)(&node),
	})
}

func (node *AccessPriv) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AccessPriv) Deparse() string {
	panic("Not Implemented")
}
