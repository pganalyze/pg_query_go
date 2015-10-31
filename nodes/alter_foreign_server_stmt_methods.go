// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterForeignServerStmt) MarshalJSON() ([]byte, error) {
	type AlterForeignServerStmtMarshalAlias AlterForeignServerStmt
	return json.Marshal(map[string]interface{}{
		"ALTERFOREIGNSERVERSTMT": (*AlterForeignServerStmtMarshalAlias)(&node),
	})
}

func (node *AlterForeignServerStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterForeignServerStmt) Deparse() string {
	panic("Not Implemented")
}
