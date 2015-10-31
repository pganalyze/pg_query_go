// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node IndexElem) MarshalJSON() ([]byte, error) {
	type IndexElemMarshalAlias IndexElem
	return json.Marshal(map[string]interface{}{
		"INDEXELEM": (*IndexElemMarshalAlias)(&node),
	})
}

func (node *IndexElem) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node IndexElem) Deparse() string {
	panic("Not Implemented")
}
