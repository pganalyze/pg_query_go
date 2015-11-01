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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
