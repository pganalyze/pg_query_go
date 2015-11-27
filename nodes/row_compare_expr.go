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
