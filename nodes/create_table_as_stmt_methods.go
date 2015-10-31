// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node CreateTableAsStmt) Deparse() string {
	panic("Not Implemented")
}
