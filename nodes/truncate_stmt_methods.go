// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node TruncateStmt) Deparse() string {
	panic("Not Implemented")
}
