// Auto-generated - DO NOT EDIT

package pg_query

type DropUserMappingStmt struct {
	Username   *string `json:"username"`   /* username or PUBLIC/CURRENT_USER */
	Servername *string `json:"servername"` /* server name */
	MissingOk  bool    `json:"missing_ok"` /* ignore missing mappings */
}
