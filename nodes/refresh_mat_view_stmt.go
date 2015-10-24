// Auto-generated - DO NOT EDIT

package pg_query

type RefreshMatViewStmt struct {
	Concurrent bool      `json:"concurrent"` /* allow concurrent access? */
	SkipData   bool      `json:"skipData"`   /* true for WITH NO DATA */
	Relation   *RangeVar `json:"relation"`   /* relation to insert into */
}
