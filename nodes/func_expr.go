// Auto-generated - DO NOT EDIT

package pg_query

type FuncExpr struct {
	Xpr            Expr `json:"xpr"`
	Funcid         Oid  `json:"funcid"`         /* PG_PROC OID of the function */
	Funcresulttype Oid  `json:"funcresulttype"` /* PG_TYPE OID of result value */
	Funcretset     bool `json:"funcretset"`     /* true if function returns set */
	Funcvariadic   bool `json:"funcvariadic"`   /* true if variadic arguments have been
	 * combined into an array last argument */
	Funcformat  CoercionForm `json:"funcformat"`  /* how to display this function call */
	Funccollid  Oid          `json:"funccollid"`  /* OID of collation of result */
	Inputcollid Oid          `json:"inputcollid"` /* OID of collation that function should use */
	Args        []Node       `json:"args"`        /* arguments to the function */
	Location    int          `json:"location"`    /* token location, or -1 if unknown */
}
