// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type MergeAppend struct {
	Plan       Plan   `json:"plan"`
	Mergeplans []Node `json:"mergeplans"`
	/* remaining fields are just like the sort-key info in struct Sort */
	NumCols       int         `json:"numCols"`       /* number of sort-key columns */
	SortColIdx    *AttrNumber `json:"sortColIdx"`    /* their indexes in the target list */
	SortOperators *Oid        `json:"sortOperators"` /* OIDs of operators to sort them by */
	Collations    *Oid        `json:"collations"`    /* OIDs of collations */
	NullsFirst    *bool       `json:"nullsFirst"`    /* NULLS FIRST/LAST directions */
}

func (node MergeAppend) MarshalJSON() ([]byte, error) {
	type MergeAppendMarshalAlias MergeAppend
	return json.Marshal(map[string]interface{}{
		"MERGEAPPEND": (*MergeAppendMarshalAlias)(&node),
	})
}

func (node *MergeAppend) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
