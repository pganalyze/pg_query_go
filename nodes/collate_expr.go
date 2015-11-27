// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CollateExpr struct {
	Xpr      Expr  `json:"xpr"`
	Arg      *Expr `json:"arg"`      /* input expression */
	CollOid  Oid   `json:"collOid"`  /* collation's OID */
	Location int   `json:"location"` /* token location, or -1 if unknown */
}

func (node CollateExpr) MarshalJSON() ([]byte, error) {
	type CollateExprMarshalAlias CollateExpr
	return json.Marshal(map[string]interface{}{
		"COLLATE": (*CollateExprMarshalAlias)(&node),
	})
}

func (node *CollateExpr) UnmarshalJSON(input []byte) (err error) {
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

	if fields["collOid"] != nil {
		err = json.Unmarshal(fields["collOid"], &node.CollOid)
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}
