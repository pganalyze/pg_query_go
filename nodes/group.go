// Auto-generated - DO NOT EDIT

package pg_query

type Group struct {
	Plan         Plan        `json:"plan"`
	NumCols      int         `json:"numCols"`      /* number of grouping columns */
	GrpColIdx    *AttrNumber `json:"grpColIdx"`    /* their indexes in the target list */
	GrpOperators *Oid        `json:"grpOperators"` /* equality operators to compare with */
}
