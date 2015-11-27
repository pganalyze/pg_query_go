// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type RowExpr struct {
	Xpr       Expr   `json:"xpr"`
	Args      []Node `json:"args"`       /* the fields */
	RowTypeid Oid    `json:"row_typeid"` /* RECORDOID or a composite type's ID */

	/*
	 * Note: we deliberately do NOT store a typmod.  Although a typmod will be
	 * associated with specific RECORD types at runtime, it will differ for
	 * different backends, and so cannot safely be stored in stored
	 * parsetrees.  We must assume typmod -1 for a RowExpr node.
	 *
	 * We don't need to store a collation either.  The result type is
	 * necessarily composite, and composite types never have a collation.
	 */
	RowFormat CoercionForm `json:"row_format"` /* how to display this node */
	Colnames  []Node       `json:"colnames"`   /* list of String, or NIL */
	Location  int          `json:"location"`   /* token location, or -1 if unknown */
}

func (node RowExpr) MarshalJSON() ([]byte, error) {
	type RowExprMarshalAlias RowExpr
	return json.Marshal(map[string]interface{}{
		"ROW": (*RowExprMarshalAlias)(&node),
	})
}

func (node *RowExpr) UnmarshalJSON(input []byte) (err error) {
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

	if fields["args"] != nil {
		node.Args, err = UnmarshalNodeArrayJSON(fields["args"])
		if err != nil {
			return
		}
	}

	if fields["row_typeid"] != nil {
		err = json.Unmarshal(fields["row_typeid"], &node.RowTypeid)
		if err != nil {
			return
		}
	}

	if fields["row_format"] != nil {
		err = json.Unmarshal(fields["row_format"], &node.RowFormat)
		if err != nil {
			return
		}
	}

	if fields["colnames"] != nil {
		node.Colnames, err = UnmarshalNodeArrayJSON(fields["colnames"])
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
