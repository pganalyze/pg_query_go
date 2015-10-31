// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterFunctionStmt) MarshalJSON() ([]byte, error) {
	type AlterFunctionStmtMarshalAlias AlterFunctionStmt
	return json.Marshal(map[string]interface{}{
		"ALTERFUNCTIONSTMT": (*AlterFunctionStmtMarshalAlias)(&node),
	})
}

func (node *AlterFunctionStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterFunctionStmt) Deparse() string {
	panic("Not Implemented")
}
