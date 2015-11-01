// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterForeignServerStmt struct {
	Servername *string `json:"servername"`  /* server name */
	Version    *string `json:"version"`     /* optional server version */
	Options    []Node  `json:"options"`     /* generic options to server */
	HasVersion bool    `json:"has_version"` /* version specified */
}

func (node AlterForeignServerStmt) MarshalJSON() ([]byte, error) {
	type AlterForeignServerStmtMarshalAlias AlterForeignServerStmt
	return json.Marshal(map[string]interface{}{
		"ALTERFOREIGNSERVERSTMT": (*AlterForeignServerStmtMarshalAlias)(&node),
	})
}

func (node *AlterForeignServerStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
