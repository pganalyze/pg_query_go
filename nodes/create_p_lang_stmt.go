// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreatePLangStmt struct {
	Replace     bool    `json:"replace"`     /* T => replace if already exists */
	Plname      *string `json:"plname"`      /* PL name */
	Plhandler   []Node  `json:"plhandler"`   /* PL call handler function (qual. name) */
	Plinline    []Node  `json:"plinline"`    /* optional inline function (qual. name) */
	Plvalidator []Node  `json:"plvalidator"` /* optional validator function (qual. name) */
	Pltrusted   bool    `json:"pltrusted"`   /* PL is trusted */
}

func (node CreatePLangStmt) MarshalJSON() ([]byte, error) {
	type CreatePLangStmtMarshalAlias CreatePLangStmt
	return json.Marshal(map[string]interface{}{
		"CREATEPLANGSTMT": (*CreatePLangStmtMarshalAlias)(&node),
	})
}

func (node *CreatePLangStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
