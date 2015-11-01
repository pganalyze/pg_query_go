// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Unique struct {
	Plan          Plan        `json:"plan"`
	NumCols       int         `json:"numCols"`       /* number of columns to check for uniqueness */
	UniqColIdx    *AttrNumber `json:"uniqColIdx"`    /* their indexes in the target list */
	UniqOperators *Oid        `json:"uniqOperators"` /* equality operators to compare with */
}

func (node Unique) MarshalJSON() ([]byte, error) {
	type UniqueMarshalAlias Unique
	return json.Marshal(map[string]interface{}{
		"UNIQUE": (*UniqueMarshalAlias)(&node),
	})
}

func (node *Unique) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
