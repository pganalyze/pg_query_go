// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Param) MarshalJSON() ([]byte, error) {
	type ParamMarshalAlias Param
	return json.Marshal(map[string]interface{}{
		"PARAM": (*ParamMarshalAlias)(&node),
	})
}

func (node *Param) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Param) Deparse() string {
	panic("Not Implemented")
}
