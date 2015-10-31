// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SetToDefault) MarshalJSON() ([]byte, error) {
	type SetToDefaultMarshalAlias SetToDefault
	return json.Marshal(map[string]interface{}{
		"SETTODEFAULT": (*SetToDefaultMarshalAlias)(&node),
	})
}

func (node *SetToDefault) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SetToDefault) Deparse() string {
	panic("Not Implemented")
}
