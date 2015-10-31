// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PlaceHolderInfo) MarshalJSON() ([]byte, error) {
	type PlaceHolderInfoMarshalAlias PlaceHolderInfo
	return json.Marshal(map[string]interface{}{
		"PLACEHOLDERINFO": (*PlaceHolderInfoMarshalAlias)(&node),
	})
}

func (node *PlaceHolderInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PlaceHolderInfo) Deparse() string {
	panic("Not Implemented")
}
