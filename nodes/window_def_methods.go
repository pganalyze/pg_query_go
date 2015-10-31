// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node WindowDef) MarshalJSON() ([]byte, error) {
	type WindowDefMarshalAlias WindowDef
	return json.Marshal(map[string]interface{}{
		"WINDOWDEF": (*WindowDefMarshalAlias)(&node),
	})
}

func (node *WindowDef) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node WindowDef) Deparse() string {
	panic("Not Implemented")
}
