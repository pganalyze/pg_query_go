// Auto-generated - DO NOT EDIT

package pg_query

type AlterDatabaseSetStmt struct {
	Dbname  *string          `json:"dbname"`  /* database name */
	Setstmt *VariableSetStmt `json:"setstmt"` /* SET or RESET subcommand */
}
