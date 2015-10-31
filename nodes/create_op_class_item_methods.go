// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateOpClassItem) MarshalJSON() ([]byte, error) {
	type CreateOpClassItemMarshalAlias CreateOpClassItem
	return json.Marshal(map[string]interface{}{
		"CREATEOPCLASSITEM": (*CreateOpClassItemMarshalAlias)(&node),
	})
}

func (node *CreateOpClassItem) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateOpClassItem) Deparse() string {
	panic("Not Implemented")
}
