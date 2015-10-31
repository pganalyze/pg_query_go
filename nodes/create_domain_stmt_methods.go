// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateDomainStmt) MarshalJSON() ([]byte, error) {
	type CreateDomainStmtMarshalAlias CreateDomainStmt
	return json.Marshal(map[string]interface{}{
		"CREATEDOMAINSTMT": (*CreateDomainStmtMarshalAlias)(&node),
	})
}

func (node *CreateDomainStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateDomainStmt) Deparse() string {
	panic("Not Implemented")
}
