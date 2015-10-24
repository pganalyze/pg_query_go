// Auto-generated - DO NOT EDIT

package pg_query

type DropdbStmt struct {
	Dbname    *string `json:"dbname"`     /* database to drop */
	MissingOk bool    `json:"missing_ok"` /* skip error if db is missing? */
}
