// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateExtensionStmt struct {
	Extname     *string `json:"extname"`
	IfNotExists bool    `json:"if_not_exists"` /* just do nothing if it already exists? */
	Options     []Node  `json:"options"`       /* List of DefElem nodes */
}

func (node CreateExtensionStmt) MarshalJSON() ([]byte, error) {
	type CreateExtensionStmtMarshalAlias CreateExtensionStmt
	return json.Marshal(map[string]interface{}{
		"CREATEEXTENSIONSTMT": (*CreateExtensionStmtMarshalAlias)(&node),
	})
}

func (node *CreateExtensionStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
