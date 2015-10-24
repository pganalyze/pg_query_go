// Auto-generated - DO NOT EDIT

package pg_query

type FunctionParameter struct {
	Name    *string               `json:"name"`    /* parameter name, or NULL if not given */
	ArgType *TypeName             `json:"argType"` /* TypeName for parameter type */
	Mode    FunctionParameterMode `json:"mode"`    /* IN/OUT/etc */
	Defexpr Node                  `json:"defexpr"` /* raw default expr, or NULL if not given */
}
