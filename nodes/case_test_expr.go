// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CaseTestExpr struct {
	Xpr       Expr  `json:"xpr"`
	TypeId    Oid   `json:"typeId"`    /* type for substituted value */
	TypeMod   int32 `json:"typeMod"`   /* typemod for substituted value */
	Collation Oid   `json:"collation"` /* collation for the substituted value */
}

func (node CaseTestExpr) MarshalJSON() ([]byte, error) {
	type CaseTestExprMarshalAlias CaseTestExpr
	return json.Marshal(map[string]interface{}{
		"CASETESTEXPR": (*CaseTestExprMarshalAlias)(&node),
	})
}

func (node *CaseTestExpr) UnmarshalJSON(input []byte) (err error) {
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

	if fields["typeId"] != nil {
		err = json.Unmarshal(fields["typeId"], &node.TypeId)
		if err != nil {
			return
		}
	}

	if fields["typeMod"] != nil {
		err = json.Unmarshal(fields["typeMod"], &node.TypeMod)
		if err != nil {
			return
		}
	}

	if fields["collation"] != nil {
		err = json.Unmarshal(fields["collation"], &node.Collation)
		if err != nil {
			return
		}
	}

	return
}
