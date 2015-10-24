// Auto-generated - DO NOT EDIT

package pg_query

type WithClause struct {
	Ctes      []Node `json:"ctes"`      /* list of CommonTableExprs */
	Recursive bool   `json:"recursive"` /* true = WITH RECURSIVE */
	Location  int    `json:"location"`  /* token location, or -1 if unknown */
}
