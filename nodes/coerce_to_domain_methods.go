// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CoerceToDomain) MarshalJSON() ([]byte, error) {
	type CoerceToDomainMarshalAlias CoerceToDomain
	return json.Marshal(map[string]interface{}{
		"COERCETODOMAIN": (*CoerceToDomainMarshalAlias)(&node),
	})
}

func (node *CoerceToDomain) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CoerceToDomain) Deparse() string {
	panic("Not Implemented")
}
