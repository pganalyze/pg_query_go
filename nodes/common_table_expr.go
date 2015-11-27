// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CommonTableExpr struct {
	Ctename       *string `json:"ctename"`       /* query name (never qualified) */
	Aliascolnames []Node  `json:"aliascolnames"` /* optional list of column names */

	/* SelectStmt/InsertStmt/etc before parse analysis, Query afterwards: */
	Ctequery Node `json:"ctequery"` /* the CTE's subquery */
	Location int  `json:"location"` /* token location, or -1 if unknown */

	/* These fields are set during parse analysis: */
	Cterecursive bool `json:"cterecursive"` /* is this CTE actually recursive? */
	Cterefcount  int  `json:"cterefcount"`  /* number of RTEs referencing this CTE
	 * (excluding internal self-references) */

	Ctecolnames      []Node `json:"ctecolnames"`      /* list of output column names */
	Ctecoltypes      []Node `json:"ctecoltypes"`      /* OID list of output column type OIDs */
	Ctecoltypmods    []Node `json:"ctecoltypmods"`    /* integer list of output column typmods */
	Ctecolcollations []Node `json:"ctecolcollations"` /* OID list of column collation OIDs */
}

func (node CommonTableExpr) MarshalJSON() ([]byte, error) {
	type CommonTableExprMarshalAlias CommonTableExpr
	return json.Marshal(map[string]interface{}{
		"COMMONTABLEEXPR": (*CommonTableExprMarshalAlias)(&node),
	})
}

func (node *CommonTableExpr) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["ctename"] != nil {
		err = json.Unmarshal(fields["ctename"], &node.Ctename)
		if err != nil {
			return
		}
	}

	if fields["aliascolnames"] != nil {
		node.Aliascolnames, err = UnmarshalNodeArrayJSON(fields["aliascolnames"])
		if err != nil {
			return
		}
	}

	if fields["ctequery"] != nil {
		node.Ctequery, err = UnmarshalNodeJSON(fields["ctequery"])
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	if fields["cterecursive"] != nil {
		err = json.Unmarshal(fields["cterecursive"], &node.Cterecursive)
		if err != nil {
			return
		}
	}

	if fields["cterefcount"] != nil {
		err = json.Unmarshal(fields["cterefcount"], &node.Cterefcount)
		if err != nil {
			return
		}
	}

	if fields["ctecolnames"] != nil {
		node.Ctecolnames, err = UnmarshalNodeArrayJSON(fields["ctecolnames"])
		if err != nil {
			return
		}
	}

	if fields["ctecoltypes"] != nil {
		node.Ctecoltypes, err = UnmarshalNodeArrayJSON(fields["ctecoltypes"])
		if err != nil {
			return
		}
	}

	if fields["ctecoltypmods"] != nil {
		node.Ctecoltypmods, err = UnmarshalNodeArrayJSON(fields["ctecoltypmods"])
		if err != nil {
			return
		}
	}

	if fields["ctecolcollations"] != nil {
		node.Ctecolcollations, err = UnmarshalNodeArrayJSON(fields["ctecolcollations"])
		if err != nil {
			return
		}
	}

	return
}
