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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
