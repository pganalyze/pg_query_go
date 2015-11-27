// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ScalarArrayOpExpr struct {
	Xpr         Expr   `json:"xpr"`
	Opno        Oid    `json:"opno"`        /* PG_OPERATOR OID of the operator */
	Opfuncid    Oid    `json:"opfuncid"`    /* PG_PROC OID of underlying function */
	UseOr       bool   `json:"useOr"`       /* true for ANY, false for ALL */
	Inputcollid Oid    `json:"inputcollid"` /* OID of collation that operator should use */
	Args        []Node `json:"args"`        /* the scalar and array operands */
	Location    int    `json:"location"`    /* token location, or -1 if unknown */
}

func (node ScalarArrayOpExpr) MarshalJSON() ([]byte, error) {
	type ScalarArrayOpExprMarshalAlias ScalarArrayOpExpr
	return json.Marshal(map[string]interface{}{
		"SCALARARRAYOPEXPR": (*ScalarArrayOpExprMarshalAlias)(&node),
	})
}

func (node *ScalarArrayOpExpr) UnmarshalJSON(input []byte) (err error) {
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

	if fields["opno"] != nil {
		err = json.Unmarshal(fields["opno"], &node.Opno)
		if err != nil {
			return
		}
	}

	if fields["opfuncid"] != nil {
		err = json.Unmarshal(fields["opfuncid"], &node.Opfuncid)
		if err != nil {
			return
		}
	}

	if fields["useOr"] != nil {
		err = json.Unmarshal(fields["useOr"], &node.UseOr)
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
