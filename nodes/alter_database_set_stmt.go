// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterDatabaseSetStmt struct {
	Dbname  *string          `json:"dbname"`  /* database name */
	Setstmt *VariableSetStmt `json:"setstmt"` /* SET or RESET subcommand */
}

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
