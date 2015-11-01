// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type FuncWithArgs struct {
	Funcname []Node `json:"funcname"` /* qualified name of function */
	Funcargs []Node `json:"funcargs"` /* list of Typename nodes */
}

func (node FuncWithArgs) MarshalJSON() ([]byte, error) {
	type FuncWithArgsMarshalAlias FuncWithArgs
	return json.Marshal(map[string]interface{}{
		"FUNCWITHARGS": (*FuncWithArgsMarshalAlias)(&node),
	})
}

func (node *FuncWithArgs) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
