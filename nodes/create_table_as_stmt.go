// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateTableAsStmt struct {
	Query        Node        `json:"query"`          /* the query (see comments above) */
	Into         *IntoClause `json:"into"`           /* destination table */
	Relkind      ObjectType  `json:"relkind"`        /* OBJECT_TABLE or OBJECT_MATVIEW */
	IsSelectInto bool        `json:"is_select_into"` /* it was written as SELECT INTO */
}

func (node CreateTableAsStmt) MarshalJSON() ([]byte, error) {
	type CreateTableAsStmtMarshalAlias CreateTableAsStmt
	return json.Marshal(map[string]interface{}{
		"CREATE TABLE AS": (*CreateTableAsStmtMarshalAlias)(&node),
	})
}

func (node *CreateTableAsStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
