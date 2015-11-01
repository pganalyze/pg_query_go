// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type TruncateStmt struct {
	Relations   []Node       `json:"relations"`    /* relations (RangeVars) to be truncated */
	RestartSeqs bool         `json:"restart_seqs"` /* restart owned sequences? */
	Behavior    DropBehavior `json:"behavior"`     /* RESTRICT or CASCADE behavior */
}

func (node TruncateStmt) MarshalJSON() ([]byte, error) {
	type TruncateStmtMarshalAlias TruncateStmt
	return json.Marshal(map[string]interface{}{
		"TRUNCATE": (*TruncateStmtMarshalAlias)(&node),
	})
}

func (node *TruncateStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
