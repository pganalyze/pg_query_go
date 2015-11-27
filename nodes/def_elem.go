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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["defnamespace"] != nil {
		err = json.Unmarshal(fields["defnamespace"], &node.Defnamespace)
		if err != nil {
			return
		}
	}

	if fields["defname"] != nil {
		err = json.Unmarshal(fields["defname"], &node.Defname)
		if err != nil {
			return
		}
	}

	if fields["arg"] != nil {
		node.Arg, err = UnmarshalNodeJSON(fields["arg"])
		if err != nil {
			return
		}
	}

	if fields["defaction"] != nil {
		err = json.Unmarshal(fields["defaction"], &node.Defaction)
		if err != nil {
			return
		}
	}

	return
}
