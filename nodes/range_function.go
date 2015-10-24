// Auto-generated - DO NOT EDIT

package pg_query

type RangeFunction struct {
	Lateral    bool   `json:"lateral"`     /* does it have LATERAL prefix? */
	Ordinality bool   `json:"ordinality"`  /* does it have WITH ORDINALITY suffix? */
	IsRowsfrom bool   `json:"is_rowsfrom"` /* is result of ROWS FROM() syntax? */
	Functions  []Node `json:"functions"`   /* per-function information, see above */
	Alias      *Alias `json:"alias"`       /* table alias & optional column aliases */
	Coldeflist []Node `json:"coldeflist"`  /* list of ColumnDef nodes to describe result
	 * of function returning RECORD */
}
