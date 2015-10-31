// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ArrayRef) MarshalJSON() ([]byte, error) {
	type ArrayRefMarshalAlias ArrayRef
	return json.Marshal(map[string]interface{}{
		"ARRAYREF": (*ArrayRefMarshalAlias)(&node),
	})
}

func (node *ArrayRef) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ArrayRef) Deparse() string {
	panic("Not Implemented")
}
