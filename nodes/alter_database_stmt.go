// Auto-generated - DO NOT EDIT

package pg_query

type AlterDatabaseStmt struct {
	Dbname  *string `json:"dbname"`  /* name of database to alter */
	Options []Node  `json:"options"` /* List of DefElem nodes */
}
