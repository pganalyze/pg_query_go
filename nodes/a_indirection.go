// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["arg"] != nil {
		node.Arg, err = UnmarshalNodeJSON(fields["arg"])
		if err != nil {
			return
		}
	}

	if fields["indirection"] != nil {
		node.Indirection, err = UnmarshalNodeArrayJSON(fields["indirection"])
		if err != nil {
			return
		}
	}

	return
}
