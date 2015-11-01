// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Sort struct {
	Plan          Plan        `json:"plan"`
	NumCols       int         `json:"numCols"`       /* number of sort-key columns */
	SortColIdx    *AttrNumber `json:"sortColIdx"`    /* their indexes in the target list */
	SortOperators *Oid        `json:"sortOperators"` /* OIDs of operators to sort them by */
	Collations    *Oid        `json:"collations"`    /* OIDs of collations */
	NullsFirst    *bool       `json:"nullsFirst"`    /* NULLS FIRST/LAST directions */
}

func (node Sort) MarshalJSON() ([]byte, error) {
	type SortMarshalAlias Sort
	return json.Marshal(map[string]interface{}{
		"SORT": (*SortMarshalAlias)(&node),
	})
}

func (node *Sort) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
