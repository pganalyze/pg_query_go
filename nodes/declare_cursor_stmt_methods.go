// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node DeclareCursorStmt) Deparse() string {
	panic("Not Implemented")
}
