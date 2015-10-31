// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node FunctionParameter) MarshalJSON() ([]byte, error) {
	type FunctionParameterMarshalAlias FunctionParameter
	return json.Marshal(map[string]interface{}{
		"FUNCTIONPARAMETER": (*FunctionParameterMarshalAlias)(&node),
	})
}

func (node *FunctionParameter) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node FunctionParameter) Deparse() string {
	panic("Not Implemented")
}
