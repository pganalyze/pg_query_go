// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ---------------
 *	 group node -
 *		Used for queries with GROUP BY (but no aggregates) specified.
 *		The input must be presorted according to the grouping columns.
 * ---------------
 */
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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["numCols"] != nil {
		err = json.Unmarshal(fields["numCols"], &node.NumCols)
		if err != nil {
			return
		}
	}

	if fields["grpColIdx"] != nil {
		err = json.Unmarshal(fields["grpColIdx"], &node.GrpColIdx)
		if err != nil {
			return
		}
	}

	if fields["grpOperators"] != nil {
		err = json.Unmarshal(fields["grpOperators"], &node.GrpOperators)
		if err != nil {
			return
		}
	}

	return
}
