// Auto-generated - DO NOT EDIT

package pg_query

type TruncateStmt struct {
	Relations   []Node       `json:"relations"`    /* relations (RangeVars) to be truncated */
	RestartSeqs bool         `json:"restart_seqs"` /* restart owned sequences? */
	Behavior    DropBehavior `json:"behavior"`     /* RESTRICT or CASCADE behavior */
}
