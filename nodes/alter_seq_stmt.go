// Auto-generated - DO NOT EDIT

package pg_query

type AlterSeqStmt struct {
	Sequence  *RangeVar `json:"sequence"` /* the sequence to alter */
	Options   []Node    `json:"options"`
	MissingOk bool      `json:"missing_ok"` /* skip error if a role is missing? */
}
