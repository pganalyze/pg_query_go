// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ReindexStmt struct {
	Kind     ObjectType `json:"kind"`      /* OBJECT_INDEX, OBJECT_TABLE, etc. */
	Relation *RangeVar  `json:"relation"`  /* Table or index to reindex */
	Name     *string    `json:"name"`      /* name of database to reindex */
	DoSystem bool       `json:"do_system"` /* include system tables in database case */
	DoUser   bool       `json:"do_user"`   /* include user tables in database case */
}

func (node ReindexStmt) MarshalJSON() ([]byte, error) {
	type ReindexStmtMarshalAlias ReindexStmt
	return json.Marshal(map[string]interface{}{
		"REINDEXSTMT": (*ReindexStmtMarshalAlias)(&node),
	})
}

func (node *ReindexStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
