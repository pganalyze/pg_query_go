// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterDomainStmt) MarshalJSON() ([]byte, error) {
	type AlterDomainStmtMarshalAlias AlterDomainStmt
	return json.Marshal(map[string]interface{}{
		"ALTERDOMAINSTMT": (*AlterDomainStmtMarshalAlias)(&node),
	})
}

func (node *AlterDomainStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterDomainStmt) Deparse() string {
	panic("Not Implemented")
}
