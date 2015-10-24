// Auto-generated - DO NOT EDIT

package pg_query

type CreateForeignServerStmt struct {
	Servername *string `json:"servername"` /* server name */
	Servertype *string `json:"servertype"` /* optional server type */
	Version    *string `json:"version"`    /* optional server version */
	Fdwname    *string `json:"fdwname"`    /* FDW name */
	Options    []Node  `json:"options"`    /* generic options to server */
}
