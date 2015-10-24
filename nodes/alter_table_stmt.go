// Auto-generated - DO NOT EDIT

package pg_query

type AlterTableStmt struct {
	Relation  *RangeVar  `json:"relation"`   /* table to work on */
	Cmds      []Node     `json:"cmds"`       /* list of subcommands */
	Relkind   ObjectType `json:"relkind"`    /* type of object */
	MissingOk bool       `json:"missing_ok"` /* skip error if table missing */
}
