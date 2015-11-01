// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DeclareCursorStmt struct {
	Portalname *string `json:"portalname"` /* name of the portal (cursor) */
	Options    int     `json:"options"`    /* bitmask of options (see above) */
	Query      Node    `json:"query"`      /* the raw SELECT query */
}

func (node DeclareCursorStmt) MarshalJSON() ([]byte, error) {
	type DeclareCursorStmtMarshalAlias DeclareCursorStmt
	return json.Marshal(map[string]interface{}{
		"DECLARECURSOR": (*DeclareCursorStmtMarshalAlias)(&node),
	})
}

func (node *DeclareCursorStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
