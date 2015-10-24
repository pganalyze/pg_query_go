// Auto-generated - DO NOT EDIT

package pg_query

type ReindexStmt struct {
	Kind     ObjectType `json:"kind"`      /* OBJECT_INDEX, OBJECT_TABLE, etc. */
	Relation *RangeVar  `json:"relation"`  /* Table or index to reindex */
	Name     *string    `json:"name"`      /* name of database to reindex */
	DoSystem bool       `json:"do_system"` /* include system tables in database case */
	DoUser   bool       `json:"do_user"`   /* include user tables in database case */
}
