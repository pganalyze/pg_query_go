// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node ParamRef) Deparse() string {
	panic("Not Implemented")
}
