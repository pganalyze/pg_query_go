// Auto-generated - DO NOT EDIT

package pg_query

type MinMaxExpr struct {
	Xpr          Expr     `json:"xpr"`
	Minmaxtype   Oid      `json:"minmaxtype"`   /* common type of arguments and result */
	Minmaxcollid Oid      `json:"minmaxcollid"` /* OID of collation of result */
	Inputcollid  Oid      `json:"inputcollid"`  /* OID of collation that function should use */
	Op           MinMaxOp `json:"op"`           /* function to execute */
	Args         []Node   `json:"args"`         /* the arguments */
	Location     int      `json:"location"`     /* token location, or -1 if unknown */
}
