// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Agg struct {
	Plan         Plan        `json:"plan"`
	Aggstrategy  AggStrategy `json:"aggstrategy"`
	NumCols      int         `json:"numCols"`      /* number of grouping columns */
	GrpColIdx    *AttrNumber `json:"grpColIdx"`    /* their indexes in the target list */
	GrpOperators *Oid        `json:"grpOperators"` /* equality operators to compare with */
	NumGroups    int64       `json:"numGroups"`    /* estimated number of groups in input */
}

func (node Agg) MarshalJSON() ([]byte, error) {
	type AggMarshalAlias Agg
	return json.Marshal(map[string]interface{}{
		"AGG": (*AggMarshalAlias)(&node),
	})
}

func (node *Agg) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
