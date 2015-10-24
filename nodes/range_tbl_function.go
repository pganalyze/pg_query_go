// Auto-generated - DO NOT EDIT

package pg_query

type RangeTblFunction struct {
	Funcexpr     Node `json:"funcexpr"`     /* expression tree for func call */
	Funccolcount int  `json:"funccolcount"` /* number of columns it contributes to RTE */
	/* These fields record the contents of a column definition list, if any: */
	Funccolnames      []Node `json:"funccolnames"`      /* column names (list of String) */
	Funccoltypes      []Node `json:"funccoltypes"`      /* OID list of column type OIDs */
	Funccoltypmods    []Node `json:"funccoltypmods"`    /* integer list of column typmods */
	Funccolcollations []Node `json:"funccolcollations"` /* OID list of column collation OIDs */
	/* This is set during planning for use by the executor: */
	Funcparams []uint32 `json:"funcparams"` /* PARAM_EXEC Param IDs affecting this func */
}
