// Auto-generated - DO NOT EDIT

package pg_query

type FuncCall struct {
	Funcname       []Node     `json:"funcname"`         /* qualified name of function */
	Args           []Node     `json:"args"`             /* the arguments (list of exprs) */
	AggOrder       []Node     `json:"agg_order"`        /* ORDER BY (list of SortBy) */
	AggFilter      Node       `json:"agg_filter"`       /* FILTER clause, if any */
	AggWithinGroup bool       `json:"agg_within_group"` /* ORDER BY appeared in WITHIN GROUP */
	AggStar        bool       `json:"agg_star"`         /* argument was really '*' */
	AggDistinct    bool       `json:"agg_distinct"`     /* arguments were labeled DISTINCT */
	FuncVariadic   bool       `json:"func_variadic"`    /* last argument was labeled VARIADIC */
	Over           *WindowDef `json:"over"`             /* OVER clause, if any */
	Location       int        `json:"location"`         /* token location, or -1 if unknown */
}
