// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ParamRef struct {
	Number   int `json:"number"`   /* the number of the parameter */
	Location int `json:"location"` /* token location, or -1 if unknown */
}

func (node ParamRef) MarshalJSON() ([]byte, error) {
	type ParamRefMarshalAlias ParamRef
	return json.Marshal(map[string]interface{}{
		"PARAMREF": (*ParamRefMarshalAlias)(&node),
	})
}

func (node *ParamRef) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
