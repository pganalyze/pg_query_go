// Auto-generated - DO NOT EDIT

package pg_query

type CreateSchemaStmt struct {
	Schemaname  *string `json:"schemaname"`    /* the name of the schema to create */
	Authid      *string `json:"authid"`        /* the owner of the created schema */
	SchemaElts  []Node  `json:"schemaElts"`    /* schema components (list of parsenodes) */
	IfNotExists bool    `json:"if_not_exists"` /* just do nothing if schema already exists? */
}
