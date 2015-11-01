// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CaseExpr struct {
	Xpr        Expr   `json:"xpr"`
	Casetype   Oid    `json:"casetype"`   /* type of expression result */
	Casecollid Oid    `json:"casecollid"` /* OID of collation, or InvalidOid if none */
	Arg        *Expr  `json:"arg"`        /* implicit equality comparison argument */
	Args       []Node `json:"args"`       /* the arguments (list of WHEN clauses) */
	Defresult  *Expr  `json:"defresult"`  /* the default result (ELSE clause) */
	Location   int    `json:"location"`   /* token location, or -1 if unknown */
}

func (node CaseExpr) MarshalJSON() ([]byte, error) {
	type CaseExprMarshalAlias CaseExpr
	return json.Marshal(map[string]interface{}{
		"CASE": (*CaseExprMarshalAlias)(&node),
	})
}

func (node *CaseExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
