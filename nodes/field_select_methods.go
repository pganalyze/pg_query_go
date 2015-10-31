// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node FieldSelect) MarshalJSON() ([]byte, error) {
	type FieldSelectMarshalAlias FieldSelect
	return json.Marshal(map[string]interface{}{
		"FIELDSELECT": (*FieldSelectMarshalAlias)(&node),
	})
}

func (node *FieldSelect) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node FieldSelect) Deparse() string {
	panic("Not Implemented")
}
