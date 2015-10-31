// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateOpClassStmt) MarshalJSON() ([]byte, error) {
	type CreateOpClassStmtMarshalAlias CreateOpClassStmt
	return json.Marshal(map[string]interface{}{
		"CREATEOPCLASSSTMT": (*CreateOpClassStmtMarshalAlias)(&node),
	})
}

func (node *CreateOpClassStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateOpClassStmt) Deparse() string {
	panic("Not Implemented")
}
