// Auto-generated - DO NOT EDIT

package pg_query

type ScalarArrayOpExpr struct {
	Xpr         Expr   `json:"xpr"`
	Opno        Oid    `json:"opno"`        /* PG_OPERATOR OID of the operator */
	Opfuncid    Oid    `json:"opfuncid"`    /* PG_PROC OID of underlying function */
	UseOr       bool   `json:"useOr"`       /* true for ANY, false for ALL */
	Inputcollid Oid    `json:"inputcollid"` /* OID of collation that operator should use */
	Args        []Node `json:"args"`        /* the scalar and array operands */
	Location    int    `json:"location"`    /* token location, or -1 if unknown */
}
