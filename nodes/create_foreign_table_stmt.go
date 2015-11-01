// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateForeignTableStmt struct {
	Base       CreateStmt `json:"base"`
	Servername *string    `json:"servername"`
	Options    []Node     `json:"options"`
}

func (node CreateForeignTableStmt) MarshalJSON() ([]byte, error) {
	type CreateForeignTableStmtMarshalAlias CreateForeignTableStmt
	return json.Marshal(map[string]interface{}{
		"CREATEFOREIGNTABLESTMT": (*CreateForeignTableStmtMarshalAlias)(&node),
	})
}

func (node *CreateForeignTableStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
