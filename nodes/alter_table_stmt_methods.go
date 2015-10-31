// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node AlterTableStmt) Deparse() string {
	panic("Not Implemented")
}
