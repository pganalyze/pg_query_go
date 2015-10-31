// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateRangeStmt) MarshalJSON() ([]byte, error) {
	type CreateRangeStmtMarshalAlias CreateRangeStmt
	return json.Marshal(map[string]interface{}{
		"CREATERANGESTMT": (*CreateRangeStmtMarshalAlias)(&node),
	})
}

func (node *CreateRangeStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateRangeStmt) Deparse() string {
	panic("Not Implemented")
}
