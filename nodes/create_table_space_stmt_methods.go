// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateTableSpaceStmt) MarshalJSON() ([]byte, error) {
	type CreateTableSpaceStmtMarshalAlias CreateTableSpaceStmt
	return json.Marshal(map[string]interface{}{
		"CREATETABLESPACESTMT": (*CreateTableSpaceStmtMarshalAlias)(&node),
	})
}

func (node *CreateTableSpaceStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateTableSpaceStmt) Deparse() string {
	panic("Not Implemented")
}
