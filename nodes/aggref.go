// Auto-generated - DO NOT EDIT

package pg_query

type Aggref struct {
	Xpr           Expr   `json:"xpr"`
	Aggfnoid      Oid    `json:"aggfnoid"`      /* pg_proc Oid of the aggregate */
	Aggtype       Oid    `json:"aggtype"`       /* type Oid of result of the aggregate */
	Aggcollid     Oid    `json:"aggcollid"`     /* OID of collation of result */
	Inputcollid   Oid    `json:"inputcollid"`   /* OID of collation that function should use */
	Aggdirectargs []Node `json:"aggdirectargs"` /* direct arguments, if an ordered-set agg */
	Args          []Node `json:"args"`          /* aggregated arguments and sort expressions */
	Aggorder      []Node `json:"aggorder"`      /* ORDER BY (list of SortGroupClause) */
	Aggdistinct   []Node `json:"aggdistinct"`   /* DISTINCT (list of SortGroupClause) */
	Aggfilter     *Expr  `json:"aggfilter"`     /* FILTER expression, if any */
	Aggstar       bool   `json:"aggstar"`       /* TRUE if argument list was really '*' */
	Aggvariadic   bool   `json:"aggvariadic"`   /* true if variadic arguments have been
	 * combined into an array last argument */
	Aggkind     byte  `json:"aggkind"`     /* aggregate kind (see pg_aggregate.h) */
	Agglevelsup Index `json:"agglevelsup"` /* > 0 if agg belongs to outer query */
	Location    int   `json:"location"`    /* token location, or -1 if unknown */
}
