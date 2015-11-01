// Auto-generated - DO NOT EDIT

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
