// Auto-generated - DO NOT EDIT

package pg_query

type WindowFunc struct {
	Xpr         Expr   `json:"xpr"`
	Winfnoid    Oid    `json:"winfnoid"`    /* pg_proc Oid of the function */
	Wintype     Oid    `json:"wintype"`     /* type Oid of result of the window function */
	Wincollid   Oid    `json:"wincollid"`   /* OID of collation of result */
	Inputcollid Oid    `json:"inputcollid"` /* OID of collation that function should use */
	Args        []Node `json:"args"`        /* arguments to the window function */
	Aggfilter   *Expr  `json:"aggfilter"`   /* FILTER expression, if any */
	Winref      Index  `json:"winref"`      /* index of associated WindowClause */
	Winstar     bool   `json:"winstar"`     /* TRUE if argument list was really '*' */
	Winagg      bool   `json:"winagg"`      /* is function a simple aggregate? */
	Location    int    `json:"location"`    /* token location, or -1 if unknown */
}
