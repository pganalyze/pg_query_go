// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type A_Indirection struct {
	Arg         Node   `json:"arg"`         /* the thing being selected from */
	Indirection []Node `json:"indirection"` /* subscripts and/or field names and/or * */
}

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
