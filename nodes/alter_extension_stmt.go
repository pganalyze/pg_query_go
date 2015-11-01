// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterExtensionStmt struct {
	Extname *string `json:"extname"`
	Options []Node  `json:"options"` /* List of DefElem nodes */
}

func (node AlterExtensionStmt) MarshalJSON() ([]byte, error) {
	type AlterExtensionStmtMarshalAlias AlterExtensionStmt
	return json.Marshal(map[string]interface{}{
		"ALTEREXTENSIONSTMT": (*AlterExtensionStmtMarshalAlias)(&node),
	})
}

func (node *AlterExtensionStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
