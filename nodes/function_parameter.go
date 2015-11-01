// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type FunctionParameter struct {
	Name    *string               `json:"name"`    /* parameter name, or NULL if not given */
	ArgType *TypeName             `json:"argType"` /* TypeName for parameter type */
	Mode    FunctionParameterMode `json:"mode"`    /* IN/OUT/etc */
	Defexpr Node                  `json:"defexpr"` /* raw default expr, or NULL if not given */
}

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
