// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node CreateForeignServerStmt) Deparse() string {
	panic("Not Implemented")
}
