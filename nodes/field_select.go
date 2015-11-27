// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type FieldSelect struct {
	Xpr        Expr       `json:"xpr"`
	Arg        *Expr      `json:"arg"`        /* input expression */
	Fieldnum   AttrNumber `json:"fieldnum"`   /* attribute number of field to extract */
	Resulttype Oid        `json:"resulttype"` /* type of the field (result type of this
	 * node) */

	Resulttypmod int32 `json:"resulttypmod"` /* output typmod (usually -1) */
	Resultcollid Oid   `json:"resultcollid"` /* OID of collation of the field */
}

func (node FieldSelect) MarshalJSON() ([]byte, error) {
	type FieldSelectMarshalAlias FieldSelect
	return json.Marshal(map[string]interface{}{
		"FIELDSELECT": (*FieldSelectMarshalAlias)(&node),
	})
}

func (node *FieldSelect) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["arg"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["arg"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Arg = &val
		}
	}

	if fields["fieldnum"] != nil {
		err = json.Unmarshal(fields["fieldnum"], &node.Fieldnum)
		if err != nil {
			return
		}
	}

	if fields["resulttype"] != nil {
		err = json.Unmarshal(fields["resulttype"], &node.Resulttype)
		if err != nil {
			return
		}
	}

	if fields["resulttypmod"] != nil {
		err = json.Unmarshal(fields["resulttypmod"], &node.Resulttypmod)
		if err != nil {
			return
		}
	}

	if fields["resultcollid"] != nil {
		err = json.Unmarshal(fields["resultcollid"], &node.Resultcollid)
		if err != nil {
			return
		}
	}

	return
}
