// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DropdbStmt struct {
	Dbname    *string `json:"dbname"`     /* database to drop */
	MissingOk bool    `json:"missing_ok"` /* skip error if db is missing? */
}

func (node DropdbStmt) MarshalJSON() ([]byte, error) {
	type DropdbStmtMarshalAlias DropdbStmt
	return json.Marshal(map[string]interface{}{
		"DROPDBSTMT": (*DropdbStmtMarshalAlias)(&node),
	})
}

func (node *DropdbStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
