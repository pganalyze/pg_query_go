// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type RowMarkClause struct {
	Rti        Index              `json:"rti"` /* range table index of target relation */
	Strength   LockClauseStrength `json:"strength"`
	NoWait     bool               `json:"noWait"`     /* NOWAIT option */
	PushedDown bool               `json:"pushedDown"` /* pushed down from higher query level? */
}

func (node RowMarkClause) MarshalJSON() ([]byte, error) {
	type RowMarkClauseMarshalAlias RowMarkClause
	return json.Marshal(map[string]interface{}{
		"ROWMARKCLAUSE": (*RowMarkClauseMarshalAlias)(&node),
	})
}

func (node *RowMarkClause) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["rti"] != nil {
		err = json.Unmarshal(fields["rti"], &node.Rti)
		if err != nil {
			return
		}
	}

	if fields["strength"] != nil {
		err = json.Unmarshal(fields["strength"], &node.Strength)
		if err != nil {
			return
		}
	}

	if fields["noWait"] != nil {
		err = json.Unmarshal(fields["noWait"], &node.NoWait)
		if err != nil {
			return
		}
	}

	if fields["pushedDown"] != nil {
		err = json.Unmarshal(fields["pushedDown"], &node.PushedDown)
		if err != nil {
			return
		}
	}

	return
}
