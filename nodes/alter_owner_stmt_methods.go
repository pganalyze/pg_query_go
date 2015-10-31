// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterOwnerStmt) MarshalJSON() ([]byte, error) {
	type AlterOwnerStmtMarshalAlias AlterOwnerStmt
	return json.Marshal(map[string]interface{}{
		"ALTEROWNERSTMT": (*AlterOwnerStmtMarshalAlias)(&node),
	})
}

func (node *AlterOwnerStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterOwnerStmt) Deparse() string {
	panic("Not Implemented")
}
