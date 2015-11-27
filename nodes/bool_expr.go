// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type BoolExpr struct {
	Xpr      Expr         `json:"xpr"`
	Boolop   BoolExprType `json:"boolop"`
	Args     []Node       `json:"args"`     /* arguments to this expression */
	Location int          `json:"location"` /* token location, or -1 if unknown */
}

func (node BoolExpr) MarshalJSON() ([]byte, error) {
	type BoolExprMarshalAlias BoolExpr
	return json.Marshal(map[string]interface{}{
		"BOOLEXPR": (*BoolExprMarshalAlias)(&node),
	})
}

func (node *BoolExpr) UnmarshalJSON(input []byte) (err error) {
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

	if fields["boolop"] != nil {
		err = json.Unmarshal(fields["boolop"], &node.Boolop)
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
