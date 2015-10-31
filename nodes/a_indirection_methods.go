// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node A_Indirection) MarshalJSON() ([]byte, error) {
	type A_IndirectionMarshalAlias A_Indirection
	return json.Marshal(map[string]interface{}{
		"A_INDIRECTION": (*A_IndirectionMarshalAlias)(&node),
	})
}

func (node *A_Indirection) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node A_Indirection) Deparse() string {
	panic("Not Implemented")
}
