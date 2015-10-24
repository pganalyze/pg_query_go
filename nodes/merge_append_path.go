// Auto-generated - DO NOT EDIT

package pg_query

type MergeAppendPath struct {
	Path        Path    `json:"path"`
	Subpaths    []Node  `json:"subpaths"`     /* list of component Paths */
	LimitTuples float64 `json:"limit_tuples"` /* hard limit on output tuples, or -1 */
}
