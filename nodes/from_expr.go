// Auto-generated - DO NOT EDIT

package pg_query

type FromExpr struct {
	Fromlist []Node `json:"fromlist"` /* List of join subtrees */
	Quals    Node   `json:"quals"`    /* qualifiers on join, if any */
}
