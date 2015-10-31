// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ParamPathInfo) MarshalJSON() ([]byte, error) {
	type ParamPathInfoMarshalAlias ParamPathInfo
	return json.Marshal(map[string]interface{}{
		"PARAMPATHINFO": (*ParamPathInfoMarshalAlias)(&node),
	})
}

func (node *ParamPathInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ParamPathInfo) Deparse() string {
	panic("Not Implemented")
}
