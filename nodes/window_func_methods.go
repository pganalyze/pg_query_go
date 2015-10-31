// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node WindowFunc) MarshalJSON() ([]byte, error) {
	type WindowFuncMarshalAlias WindowFunc
	return json.Marshal(map[string]interface{}{
		"WINDOWFUNC": (*WindowFuncMarshalAlias)(&node),
	})
}

func (node *WindowFunc) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node WindowFunc) Deparse() string {
	panic("Not Implemented")
}
