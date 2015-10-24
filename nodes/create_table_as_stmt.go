// Auto-generated - DO NOT EDIT

package pg_query

type CreateTableAsStmt struct {
	Query        Node        `json:"query"`          /* the query (see comments above) */
	Into         *IntoClause `json:"into"`           /* destination table */
	Relkind      ObjectType  `json:"relkind"`        /* OBJECT_TABLE or OBJECT_MATVIEW */
	IsSelectInto bool        `json:"is_select_into"` /* it was written as SELECT INTO */
}
