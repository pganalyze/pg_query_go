// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Var) MarshalJSON() ([]byte, error) {
	type VarMarshalAlias Var
	return json.Marshal(map[string]interface{}{
		"VAR": (*VarMarshalAlias)(&node),
	})
}

func (node *Var) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Var) Deparse() string {
	panic("Not Implemented")
}
