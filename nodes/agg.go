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

	if fields["aggstrategy"] != nil {
		err = json.Unmarshal(fields["aggstrategy"], &node.Aggstrategy)
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

	if fields["numGroups"] != nil {
		err = json.Unmarshal(fields["numGroups"], &node.NumGroups)
		if err != nil {
			return
		}
	}

	return
}
