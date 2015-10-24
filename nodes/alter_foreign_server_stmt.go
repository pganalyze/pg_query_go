// Auto-generated - DO NOT EDIT

package pg_query

type AlterForeignServerStmt struct {
	Servername *string `json:"servername"`  /* server name */
	Version    *string `json:"version"`     /* optional server version */
	Options    []Node  `json:"options"`     /* generic options to server */
	HasVersion bool    `json:"has_version"` /* version specified */
}
