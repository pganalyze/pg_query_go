// Auto-generated - DO NOT EDIT

package pg_query

type AlterRoleSetStmt struct {
	Role     *string          `json:"role"`     /* role name */
	Database *string          `json:"database"` /* database name, or NULL */
	Setstmt  *VariableSetStmt `json:"setstmt"`  /* SET or RESET subcommand */
}
