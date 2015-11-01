// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterDatabaseStmt struct {
	Dbname  *string `json:"dbname"`  /* name of database to alter */
	Options []Node  `json:"options"` /* List of DefElem nodes */
}

func (node AlterDatabaseStmt) MarshalJSON() ([]byte, error) {
	type AlterDatabaseStmtMarshalAlias AlterDatabaseStmt
	return json.Marshal(map[string]interface{}{
		"ALTERDATABASESTMT": (*AlterDatabaseStmtMarshalAlias)(&node),
	})
}

func (node *AlterDatabaseStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
