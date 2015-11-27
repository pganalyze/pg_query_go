// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * RowCompareExpr - row-wise comparison, such as (a, b) <= (1, 2)
 *
 * We support row comparison for any operator that can be determined to
 * act like =, <>, <, <=, >, or >= (we determine this by looking for the
 * operator in btree opfamilies).  Note that the same operator name might
 * map to a different operator for each pair of row elements, since the
 * element datatypes can vary.
 *
 * A RowCompareExpr node is only generated for the < <= > >= cases;
 * the = and <> cases are translated to simple AND or OR combinations
 * of the pairwise comparisons.  However, we include = and <> in the
 * RowCompareType enum for the convenience of parser logic.
 */
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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["rctype"] != nil {
		err = json.Unmarshal(fields["rctype"], &node.Rctype)
		if err != nil {
			return
		}
	}

	if fields["opnos"] != nil {
		node.Opnos, err = UnmarshalNodeArrayJSON(fields["opnos"])
		if err != nil {
			return
		}
	}

	if fields["opfamilies"] != nil {
		node.Opfamilies, err = UnmarshalNodeArrayJSON(fields["opfamilies"])
		if err != nil {
			return
		}
	}

	if fields["inputcollids"] != nil {
		node.Inputcollids, err = UnmarshalNodeArrayJSON(fields["inputcollids"])
		if err != nil {
			return
		}
	}

	if fields["largs"] != nil {
		node.Largs, err = UnmarshalNodeArrayJSON(fields["largs"])
		if err != nil {
			return
		}
	}

	if fields["rargs"] != nil {
		node.Rargs, err = UnmarshalNodeArrayJSON(fields["rargs"])
		if err != nil {
			return
		}
	}

	return
}
