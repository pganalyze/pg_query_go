// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateSchemaStmt) MarshalJSON() ([]byte, error) {
	type CreateSchemaStmtMarshalAlias CreateSchemaStmt
	return json.Marshal(map[string]interface{}{
		"CREATE SCHEMA": (*CreateSchemaStmtMarshalAlias)(&node),
	})
}

func (node *CreateSchemaStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateSchemaStmt) Deparse() string {
	panic("Not Implemented")
}
