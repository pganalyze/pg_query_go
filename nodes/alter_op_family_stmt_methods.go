// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterOpFamilyStmt) MarshalJSON() ([]byte, error) {
	type AlterOpFamilyStmtMarshalAlias AlterOpFamilyStmt
	return json.Marshal(map[string]interface{}{
		"ALTEROPFAMILYSTMT": (*AlterOpFamilyStmtMarshalAlias)(&node),
	})
}

func (node *AlterOpFamilyStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterOpFamilyStmt) Deparse() string {
	panic("Not Implemented")
}
