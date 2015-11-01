// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RecursiveUnion struct {
	Plan    Plan `json:"plan"`
	WtParam int  `json:"wtParam"` /* ID of Param representing work table */
	/* Remaining fields are zero/null in UNION ALL case */
	NumCols int `json:"numCols"` /* number of columns to check for
	 * duplicate-ness */
	DupColIdx    *AttrNumber `json:"dupColIdx"`    /* their indexes in the target list */
	DupOperators *Oid        `json:"dupOperators"` /* equality operators to compare with */
	NumGroups    int64       `json:"numGroups"`    /* estimated number of groups in input */
}

func (node RecursiveUnion) MarshalJSON() ([]byte, error) {
	type RecursiveUnionMarshalAlias RecursiveUnion
	return json.Marshal(map[string]interface{}{
		"RECURSIVEUNION": (*RecursiveUnionMarshalAlias)(&node),
	})
}

func (node *RecursiveUnion) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
