// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateCastStmt) MarshalJSON() ([]byte, error) {
	type CreateCastStmtMarshalAlias CreateCastStmt
	return json.Marshal(map[string]interface{}{
		"CREATECASTSTMT": (*CreateCastStmtMarshalAlias)(&node),
	})
}

func (node *CreateCastStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateCastStmt) Deparse() string {
	panic("Not Implemented")
}
