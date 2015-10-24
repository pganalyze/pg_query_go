// Auto-generated - DO NOT EDIT

package pg_query

type MergeJoin struct {
	Join         Join   `json:"join"`
	Mergeclauses []Node `json:"mergeclauses"` /* mergeclauses as expression trees */
	/* these are arrays, but have the same length as the mergeclauses list: */
	MergeFamilies   *Oid  `json:"mergeFamilies"`   /* per-clause OIDs of btree opfamilies */
	MergeCollations *Oid  `json:"mergeCollations"` /* per-clause OIDs of collations */
	MergeStrategies *int  `json:"mergeStrategies"` /* per-clause ordering (ASC or DESC) */
	MergeNullsFirst *bool `json:"mergeNullsFirst"` /* per-clause nulls ordering */
}
