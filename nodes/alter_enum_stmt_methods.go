// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterEnumStmt) MarshalJSON() ([]byte, error) {
	type AlterEnumStmtMarshalAlias AlterEnumStmt
	return json.Marshal(map[string]interface{}{
		"ALTERENUMSTMT": (*AlterEnumStmtMarshalAlias)(&node),
	})
}

func (node *AlterEnumStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterEnumStmt) Deparse() string {
	panic("Not Implemented")
}
