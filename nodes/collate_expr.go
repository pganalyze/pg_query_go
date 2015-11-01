// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CollateExpr struct {
	Xpr      Expr  `json:"xpr"`
	Arg      *Expr `json:"arg"`      /* input expression */
	CollOid  Oid   `json:"collOid"`  /* collation's OID */
	Location int   `json:"location"` /* token location, or -1 if unknown */
}

func (node CollateExpr) MarshalJSON() ([]byte, error) {
	type CollateExprMarshalAlias CollateExpr
	return json.Marshal(map[string]interface{}{
		"COLLATE": (*CollateExprMarshalAlias)(&node),
	})
}

func (node *CollateExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
