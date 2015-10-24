// Auto-generated - DO NOT EDIT

package pg_query

type RowCompareExpr struct {
	Xpr          Expr           `json:"xpr"`
	Rctype       RowCompareType `json:"rctype"`       /* LT LE GE or GT, never EQ or NE */
	Opnos        []Node         `json:"opnos"`        /* OID list of pairwise comparison ops */
	Opfamilies   []Node         `json:"opfamilies"`   /* OID list of containing operator families */
	Inputcollids []Node         `json:"inputcollids"` /* OID list of collations for comparisons */
	Largs        []Node         `json:"largs"`        /* the left-hand input arguments */
	Rargs        []Node         `json:"rargs"`        /* the right-hand input arguments */
}
