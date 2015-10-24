// Auto-generated - DO NOT EDIT

package pg_query

type CollateClause struct {
	Arg      Node   `json:"arg"`      /* input expression */
	Collname []Node `json:"collname"` /* possibly-qualified collation name */
	Location int    `json:"location"` /* token location, or -1 if unknown */
}
