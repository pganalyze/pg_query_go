// Auto-generated - DO NOT EDIT

package pg_query

type SortGroupClause struct {
	TleSortGroupRef Index `json:"tleSortGroupRef"` /* reference into targetlist */
	Eqop            Oid   `json:"eqop"`            /* the equality operator ('=' op) */
	Sortop          Oid   `json:"sortop"`          /* the ordering operator ('<' op), or 0 */
	NullsFirst      bool  `json:"nulls_first"`     /* do NULLs come before normal values? */
	Hashable        bool  `json:"hashable"`        /* can eqop be implemented by hashing? */
}
