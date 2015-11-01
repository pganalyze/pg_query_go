// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Group struct {
	Plan         Plan        `json:"plan"`
	NumCols      int         `json:"numCols"`      /* number of grouping columns */
	GrpColIdx    *AttrNumber `json:"grpColIdx"`    /* their indexes in the target list */
	GrpOperators *Oid        `json:"grpOperators"` /* equality operators to compare with */
}

func (node Group) MarshalJSON() ([]byte, error) {
	type GroupMarshalAlias Group
	return json.Marshal(map[string]interface{}{
		"GROUP": (*GroupMarshalAlias)(&node),
	})
}

func (node *Group) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
