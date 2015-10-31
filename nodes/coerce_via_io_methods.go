// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CoerceViaIO) MarshalJSON() ([]byte, error) {
	type CoerceViaIOMarshalAlias CoerceViaIO
	return json.Marshal(map[string]interface{}{
		"COERCEVIAIO": (*CoerceViaIOMarshalAlias)(&node),
	})
}

func (node *CoerceViaIO) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CoerceViaIO) Deparse() string {
	panic("Not Implemented")
}
