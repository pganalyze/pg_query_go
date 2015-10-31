// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterDatabaseSetStmt) MarshalJSON() ([]byte, error) {
	type AlterDatabaseSetStmtMarshalAlias AlterDatabaseSetStmt
	return json.Marshal(map[string]interface{}{
		"ALTERDATABASESETSTMT": (*AlterDatabaseSetStmtMarshalAlias)(&node),
	})
}

func (node *AlterDatabaseSetStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterDatabaseSetStmt) Deparse() string {
	panic("Not Implemented")
}
