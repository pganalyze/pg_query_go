// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterTableStmt struct {
	Relation  *RangeVar  `json:"relation"`   /* table to work on */
	Cmds      []Node     `json:"cmds"`       /* list of subcommands */
	Relkind   ObjectType `json:"relkind"`    /* type of object */
	MissingOk bool       `json:"missing_ok"` /* skip error if table missing */
}

func (node AlterTableStmt) MarshalJSON() ([]byte, error) {
	type AlterTableStmtMarshalAlias AlterTableStmt
	return json.Marshal(map[string]interface{}{
		"ALTER TABLE": (*AlterTableStmtMarshalAlias)(&node),
	})
}

func (node *AlterTableStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
