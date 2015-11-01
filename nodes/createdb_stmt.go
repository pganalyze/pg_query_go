// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreatedbStmt struct {
	Dbname  *string `json:"dbname"`  /* name of database to create */
	Options []Node  `json:"options"` /* List of DefElem nodes */
}

func (node CreatedbStmt) MarshalJSON() ([]byte, error) {
	type CreatedbStmtMarshalAlias CreatedbStmt
	return json.Marshal(map[string]interface{}{
		"CREATEDBSTMT": (*CreatedbStmtMarshalAlias)(&node),
	})
}

func (node *CreatedbStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
