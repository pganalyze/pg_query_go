// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateExtensionStmt) MarshalJSON() ([]byte, error) {
	type CreateExtensionStmtMarshalAlias CreateExtensionStmt
	return json.Marshal(map[string]interface{}{
		"CREATEEXTENSIONSTMT": (*CreateExtensionStmtMarshalAlias)(&node),
	})
}

func (node *CreateExtensionStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateExtensionStmt) Deparse() string {
	panic("Not Implemented")
}
