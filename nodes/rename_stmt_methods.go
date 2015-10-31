// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RenameStmt) MarshalJSON() ([]byte, error) {
	type RenameStmtMarshalAlias RenameStmt
	return json.Marshal(map[string]interface{}{
		"RENAMESTMT": (*RenameStmtMarshalAlias)(&node),
	})
}

func (node *RenameStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RenameStmt) Deparse() string {
	panic("Not Implemented")
}
