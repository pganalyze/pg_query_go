// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type TypeCast struct {
	Arg      Node      `json:"arg"`      /* the expression being casted */
	TypeName *TypeName `json:"typeName"` /* the target type */
	Location int       `json:"location"` /* token location, or -1 if unknown */
}

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
