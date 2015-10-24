// Auto-generated - DO NOT EDIT

package pg_query

type UniquePath struct {
	Path        Path             `json:"path"`
	Subpath     *Path            `json:"subpath"`
	Umethod     UniquePathMethod `json:"umethod"`
	InOperators []Node           `json:"in_operators"` /* equality operators of the IN clause */
	UniqExprs   []Node           `json:"uniq_exprs"`   /* expressions to be made unique */
}
