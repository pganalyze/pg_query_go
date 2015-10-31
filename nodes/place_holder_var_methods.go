// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PlaceHolderVar) MarshalJSON() ([]byte, error) {
	type PlaceHolderVarMarshalAlias PlaceHolderVar
	return json.Marshal(map[string]interface{}{
		"PLACEHOLDERVAR": (*PlaceHolderVarMarshalAlias)(&node),
	})
}

func (node *PlaceHolderVar) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PlaceHolderVar) Deparse() string {
	panic("Not Implemented")
}
