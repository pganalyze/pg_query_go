// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RowCompareExpr struct {
	Xpr          Expr           `json:"xpr"`
	Rctype       RowCompareType `json:"rctype"`       /* LT LE GE or GT, never EQ or NE */
	Opnos        []Node         `json:"opnos"`        /* OID list of pairwise comparison ops */
	Opfamilies   []Node         `json:"opfamilies"`   /* OID list of containing operator families */
	Inputcollids []Node         `json:"inputcollids"` /* OID list of collations for comparisons */
	Largs        []Node         `json:"largs"`        /* the left-hand input arguments */
	Rargs        []Node         `json:"rargs"`        /* the right-hand input arguments */
}

func (node RowCompareExpr) MarshalJSON() ([]byte, error) {
	type RowCompareExprMarshalAlias RowCompareExpr
	return json.Marshal(map[string]interface{}{
		"ROWCOMPARE": (*RowCompareExprMarshalAlias)(&node),
	})
}

func (node *RowCompareExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
