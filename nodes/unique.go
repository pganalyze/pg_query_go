// Auto-generated - DO NOT EDIT

package pg_query

type Unique struct {
	Plan          Plan        `json:"plan"`
	NumCols       int         `json:"numCols"`       /* number of columns to check for uniqueness */
	UniqColIdx    *AttrNumber `json:"uniqColIdx"`    /* their indexes in the target list */
	UniqOperators *Oid        `json:"uniqOperators"` /* equality operators to compare with */
}
