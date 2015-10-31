// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateOpFamilyStmt) MarshalJSON() ([]byte, error) {
	type CreateOpFamilyStmtMarshalAlias CreateOpFamilyStmt
	return json.Marshal(map[string]interface{}{
		"CREATEOPFAMILYSTMT": (*CreateOpFamilyStmtMarshalAlias)(&node),
	})
}

func (node *CreateOpFamilyStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateOpFamilyStmt) Deparse() string {
	panic("Not Implemented")
}
