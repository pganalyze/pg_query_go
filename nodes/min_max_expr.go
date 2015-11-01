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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
