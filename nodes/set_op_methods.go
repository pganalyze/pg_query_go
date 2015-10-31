// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SetOp) MarshalJSON() ([]byte, error) {
	type SetOpMarshalAlias SetOp
	return json.Marshal(map[string]interface{}{
		"SETOP": (*SetOpMarshalAlias)(&node),
	})
}

func (node *SetOp) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SetOp) Deparse() string {
	panic("Not Implemented")
}
