// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateForeignTableStmt) MarshalJSON() ([]byte, error) {
	type CreateForeignTableStmtMarshalAlias CreateForeignTableStmt
	return json.Marshal(map[string]interface{}{
		"CREATEFOREIGNTABLESTMT": (*CreateForeignTableStmtMarshalAlias)(&node),
	})
}

func (node *CreateForeignTableStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateForeignTableStmt) Deparse() string {
	panic("Not Implemented")
}
