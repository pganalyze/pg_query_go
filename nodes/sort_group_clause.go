// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type SortGroupClause struct {
	TleSortGroupRef Index `json:"tleSortGroupRef"` /* reference into targetlist */
	Eqop            Oid   `json:"eqop"`            /* the equality operator ('=' op) */
	Sortop          Oid   `json:"sortop"`          /* the ordering operator ('<' op), or 0 */
	NullsFirst      bool  `json:"nulls_first"`     /* do NULLs come before normal values? */
	Hashable        bool  `json:"hashable"`        /* can eqop be implemented by hashing? */
}

func (node SortGroupClause) MarshalJSON() ([]byte, error) {
	type SortGroupClauseMarshalAlias SortGroupClause
	return json.Marshal(map[string]interface{}{
		"SORTGROUPCLAUSE": (*SortGroupClauseMarshalAlias)(&node),
	})
}

func (node *SortGroupClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
