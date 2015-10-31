// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node FieldStore) MarshalJSON() ([]byte, error) {
	type FieldStoreMarshalAlias FieldStore
	return json.Marshal(map[string]interface{}{
		"FIELDSTORE": (*FieldStoreMarshalAlias)(&node),
	})
}

func (node *FieldStore) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node FieldStore) Deparse() string {
	panic("Not Implemented")
}
