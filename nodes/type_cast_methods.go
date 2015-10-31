// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node TypeCast) MarshalJSON() ([]byte, error) {
	type TypeCastMarshalAlias TypeCast
	return json.Marshal(map[string]interface{}{
		"TYPECAST": (*TypeCastMarshalAlias)(&node),
	})
}

func (node *TypeCast) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node TypeCast) Deparse() string {
	panic("Not Implemented")
}
