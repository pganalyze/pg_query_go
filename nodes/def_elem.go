// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DefElem struct {
	Defnamespace *string       `json:"defnamespace"` /* NULL if unqualified name */
	Defname      *string       `json:"defname"`
	Arg          Node          `json:"arg"`       /* a (Value *) or a (TypeName *) */
	Defaction    DefElemAction `json:"defaction"` /* unspecified action, or SET/ADD/DROP */
}

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
