// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node CreatePLangStmt) Deparse() string {
	panic("Not Implemented")
}
