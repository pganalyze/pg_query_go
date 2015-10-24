// Auto-generated - DO NOT EDIT

package pg_query

type RangeSubselect struct {
	Lateral  bool   `json:"lateral"`  /* does it have LATERAL prefix? */
	Subquery Node   `json:"subquery"` /* the untransformed sub-select clause */
	Alias    *Alias `json:"alias"`    /* table alias & optional column aliases */
}
