// Auto-generated - DO NOT EDIT

package pg_query

type ClusterStmt struct {
	Relation  *RangeVar `json:"relation"`  /* relation being indexed, or NULL if all */
	Indexname *string   `json:"indexname"` /* original index defined */
	Verbose   bool      `json:"verbose"`   /* print progress info */
}
