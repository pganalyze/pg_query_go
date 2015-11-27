// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type SetOp struct {
	Plan     Plan          `json:"plan"`
	Cmd      SetOpCmd      `json:"cmd"`      /* what to do */
	Strategy SetOpStrategy `json:"strategy"` /* how to do it */
	NumCols  int           `json:"numCols"`  /* number of columns to check for
	 * duplicate-ness */

	DupColIdx    *AttrNumber `json:"dupColIdx"`    /* their indexes in the target list */
	DupOperators *Oid        `json:"dupOperators"` /* equality operators to compare with */
	FlagColIdx   AttrNumber  `json:"flagColIdx"`   /* where is the flag column, if any */
	FirstFlag    int         `json:"firstFlag"`    /* flag value for first input relation */
	NumGroups    int64       `json:"numGroups"`    /* estimated number of groups in input */
}

func (node SetOp) MarshalJSON() ([]byte, error) {
	type SetOpMarshalAlias SetOp
	return json.Marshal(map[string]interface{}{
		"SETOP": (*SetOpMarshalAlias)(&node),
	})
}

func (node *SetOp) UnmarshalJSON(input []byte) (err error) {
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

	if fields["cmd"] != nil {
		err = json.Unmarshal(fields["cmd"], &node.Cmd)
		if err != nil {
			return
		}
	}

	if fields["strategy"] != nil {
		err = json.Unmarshal(fields["strategy"], &node.Strategy)
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

	if fields["dupColIdx"] != nil {
		err = json.Unmarshal(fields["dupColIdx"], &node.DupColIdx)
		if err != nil {
			return
		}
	}

	if fields["dupOperators"] != nil {
		err = json.Unmarshal(fields["dupOperators"], &node.DupOperators)
		if err != nil {
			return
		}
	}

	if fields["flagColIdx"] != nil {
		err = json.Unmarshal(fields["flagColIdx"], &node.FlagColIdx)
		if err != nil {
			return
		}
	}

	if fields["firstFlag"] != nil {
		err = json.Unmarshal(fields["firstFlag"], &node.FirstFlag)
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
