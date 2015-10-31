// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateUserMappingStmt) MarshalJSON() ([]byte, error) {
	type CreateUserMappingStmtMarshalAlias CreateUserMappingStmt
	return json.Marshal(map[string]interface{}{
		"CREATEUSERMAPPINGSTMT": (*CreateUserMappingStmtMarshalAlias)(&node),
	})
}

func (node *CreateUserMappingStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateUserMappingStmt) Deparse() string {
	panic("Not Implemented")
}
