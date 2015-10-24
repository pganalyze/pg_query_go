// Auto-generated - DO NOT EDIT

package pg_query

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
