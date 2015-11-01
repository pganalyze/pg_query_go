// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateSchemaStmt struct {
	Schemaname  *string `json:"schemaname"`    /* the name of the schema to create */
	Authid      *string `json:"authid"`        /* the owner of the created schema */
	SchemaElts  []Node  `json:"schemaElts"`    /* schema components (list of parsenodes) */
	IfNotExists bool    `json:"if_not_exists"` /* just do nothing if schema already exists? */
}

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
