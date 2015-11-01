// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterRoleSetStmt struct {
	Role     *string          `json:"role"`     /* role name */
	Database *string          `json:"database"` /* database name, or NULL */
	Setstmt  *VariableSetStmt `json:"setstmt"`  /* SET or RESET subcommand */
}

func (node AlterRoleSetStmt) MarshalJSON() ([]byte, error) {
	type AlterRoleSetStmtMarshalAlias AlterRoleSetStmt
	return json.Marshal(map[string]interface{}{
		"ALTERROLESETSTMT": (*AlterRoleSetStmtMarshalAlias)(&node),
	})
}

func (node *AlterRoleSetStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
