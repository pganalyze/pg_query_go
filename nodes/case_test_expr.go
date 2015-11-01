// Auto-generated - DO NOT EDIT

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
