// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type OpExpr struct {
	Xpr          Expr   `json:"xpr"`
	Opno         Oid    `json:"opno"`         /* PG_OPERATOR OID of the operator */
	Opfuncid     Oid    `json:"opfuncid"`     /* PG_PROC OID of underlying function */
	Opresulttype Oid    `json:"opresulttype"` /* PG_TYPE OID of result value */
	Opretset     bool   `json:"opretset"`     /* true if operator returns set */
	Opcollid     Oid    `json:"opcollid"`     /* OID of collation of result */
	Inputcollid  Oid    `json:"inputcollid"`  /* OID of collation that operator should use */
	Args         []Node `json:"args"`         /* arguments to the operator (1 or 2) */
	Location     int    `json:"location"`     /* token location, or -1 if unknown */
}

func (node OpExpr) MarshalJSON() ([]byte, error) {
	type OpExprMarshalAlias OpExpr
	return json.Marshal(map[string]interface{}{
		"OPEXPR": (*OpExprMarshalAlias)(&node),
	})
}

func (node *OpExpr) UnmarshalJSON(input []byte) (err error) {
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

	if fields["opresulttype"] != nil {
		err = json.Unmarshal(fields["opresulttype"], &node.Opresulttype)
		if err != nil {
			return
		}
	}

	if fields["opretset"] != nil {
		err = json.Unmarshal(fields["opretset"], &node.Opretset)
		if err != nil {
			return
		}
	}

	if fields["opcollid"] != nil {
		err = json.Unmarshal(fields["opcollid"], &node.Opcollid)
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
