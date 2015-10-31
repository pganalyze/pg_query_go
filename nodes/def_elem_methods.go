// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DefElem) MarshalJSON() ([]byte, error) {
	type DefElemMarshalAlias DefElem
	return json.Marshal(map[string]interface{}{
		"DEFELEM": (*DefElemMarshalAlias)(&node),
	})
}

func (node *DefElem) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DefElem) Deparse() string {
	panic("Not Implemented")
}
