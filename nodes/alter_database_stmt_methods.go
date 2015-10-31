// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterDatabaseStmt) MarshalJSON() ([]byte, error) {
	type AlterDatabaseStmtMarshalAlias AlterDatabaseStmt
	return json.Marshal(map[string]interface{}{
		"ALTERDATABASESTMT": (*AlterDatabaseStmtMarshalAlias)(&node),
	})
}

func (node *AlterDatabaseStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterDatabaseStmt) Deparse() string {
	panic("Not Implemented")
}
