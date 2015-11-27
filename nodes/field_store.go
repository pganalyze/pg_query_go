// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type FieldStore struct {
	Xpr        Expr   `json:"xpr"`
	Arg        *Expr  `json:"arg"`        /* input tuple value */
	Newvals    []Node `json:"newvals"`    /* new value(s) for field(s) */
	Fieldnums  []Node `json:"fieldnums"`  /* integer list of field attnums */
	Resulttype Oid    `json:"resulttype"` /* type of result (same as type of arg) */

	/* Like RowExpr, we deliberately omit a typmod and collation here */
}

func (node FieldStore) MarshalJSON() ([]byte, error) {
	type FieldStoreMarshalAlias FieldStore
	return json.Marshal(map[string]interface{}{
		"FIELDSTORE": (*FieldStoreMarshalAlias)(&node),
	})
}

func (node *FieldStore) UnmarshalJSON(input []byte) (err error) {
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

	if fields["newvals"] != nil {
		node.Newvals, err = UnmarshalNodeArrayJSON(fields["newvals"])
		if err != nil {
			return
		}
	}

	if fields["fieldnums"] != nil {
		node.Fieldnums, err = UnmarshalNodeArrayJSON(fields["fieldnums"])
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

	return
}
