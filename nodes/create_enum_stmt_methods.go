// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateEnumStmt) MarshalJSON() ([]byte, error) {
	type CreateEnumStmtMarshalAlias CreateEnumStmt
	return json.Marshal(map[string]interface{}{
		"CREATEENUMSTMT": (*CreateEnumStmtMarshalAlias)(&node),
	})
}

func (node *CreateEnumStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateEnumStmt) Deparse() string {
	panic("Not Implemented")
}
