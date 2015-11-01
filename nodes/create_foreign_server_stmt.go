// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateForeignServerStmt struct {
	Servername *string `json:"servername"` /* server name */
	Servertype *string `json:"servertype"` /* optional server type */
	Version    *string `json:"version"`    /* optional server version */
	Fdwname    *string `json:"fdwname"`    /* FDW name */
	Options    []Node  `json:"options"`    /* generic options to server */
}

func (node CreateForeignServerStmt) MarshalJSON() ([]byte, error) {
	type CreateForeignServerStmtMarshalAlias CreateForeignServerStmt
	return json.Marshal(map[string]interface{}{
		"CREATEFOREIGNSERVERSTMT": (*CreateForeignServerStmtMarshalAlias)(&node),
	})
}

func (node *CreateForeignServerStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
