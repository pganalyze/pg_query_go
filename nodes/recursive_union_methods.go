// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RecursiveUnion) MarshalJSON() ([]byte, error) {
	type RecursiveUnionMarshalAlias RecursiveUnion
	return json.Marshal(map[string]interface{}{
		"RECURSIVEUNION": (*RecursiveUnionMarshalAlias)(&node),
	})
}

func (node *RecursiveUnion) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RecursiveUnion) Deparse() string {
	panic("Not Implemented")
}
