// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type MinMaxExpr struct {
	Xpr          Expr     `json:"xpr"`
	Minmaxtype   Oid      `json:"minmaxtype"`   /* common type of arguments and result */
	Minmaxcollid Oid      `json:"minmaxcollid"` /* OID of collation of result */
	Inputcollid  Oid      `json:"inputcollid"`  /* OID of collation that function should use */
	Op           MinMaxOp `json:"op"`           /* function to execute */
	Args         []Node   `json:"args"`         /* the arguments */
	Location     int      `json:"location"`     /* token location, or -1 if unknown */
}

func (node MinMaxExpr) MarshalJSON() ([]byte, error) {
	type MinMaxExprMarshalAlias MinMaxExpr
	return json.Marshal(map[string]interface{}{
		"MINMAX": (*MinMaxExprMarshalAlias)(&node),
	})
}

func (node *MinMaxExpr) UnmarshalJSON(input []byte) (err error) {
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

	if fields["minmaxtype"] != nil {
		err = json.Unmarshal(fields["minmaxtype"], &node.Minmaxtype)
		if err != nil {
			return
		}
	}

	if fields["minmaxcollid"] != nil {
		err = json.Unmarshal(fields["minmaxcollid"], &node.Minmaxcollid)
		if err != nil {
			return
		}
	}

	if fields["inputcollid"] != nil {
		err = json.Unmarshal(fields["inputcollid"], &node.Inputcollid)
		if err != nil {
			return
		}
	}

	if fields["op"] != nil {
		err = json.Unmarshal(fields["op"], &node.Op)
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

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}
