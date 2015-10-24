// Auto-generated - DO NOT EDIT

package pg_query

type Agg struct {
	Plan         Plan        `json:"plan"`
	Aggstrategy  AggStrategy `json:"aggstrategy"`
	NumCols      int         `json:"numCols"`      /* number of grouping columns */
	GrpColIdx    *AttrNumber `json:"grpColIdx"`    /* their indexes in the target list */
	GrpOperators *Oid        `json:"grpOperators"` /* equality operators to compare with */
	NumGroups    int64       `json:"numGroups"`    /* estimated number of groups in input */
}
