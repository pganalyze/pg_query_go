// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CoalesceExpr struct {
	Xpr            Expr   `json:"xpr"`
	Coalescetype   Oid    `json:"coalescetype"`   /* type of expression result */
	Coalescecollid Oid    `json:"coalescecollid"` /* OID of collation, or InvalidOid if none */
	Args           []Node `json:"args"`           /* the arguments */
	Location       int    `json:"location"`       /* token location, or -1 if unknown */
}

func (node CoalesceExpr) MarshalJSON() ([]byte, error) {
	type CoalesceExprMarshalAlias CoalesceExpr
	return json.Marshal(map[string]interface{}{
		"COALESCE": (*CoalesceExprMarshalAlias)(&node),
	})
}

func (node *CoalesceExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
