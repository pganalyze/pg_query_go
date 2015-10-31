// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterTableMoveAllStmt) MarshalJSON() ([]byte, error) {
	type AlterTableMoveAllStmtMarshalAlias AlterTableMoveAllStmt
	return json.Marshal(map[string]interface{}{
		"ALTERTABLEMOVEALLSTMT": (*AlterTableMoveAllStmtMarshalAlias)(&node),
	})
}

func (node *AlterTableMoveAllStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterTableMoveAllStmt) Deparse() string {
	panic("Not Implemented")
}
