// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterUserMappingStmt) MarshalJSON() ([]byte, error) {
	type AlterUserMappingStmtMarshalAlias AlterUserMappingStmt
	return json.Marshal(map[string]interface{}{
		"ALTERUSERMAPPINGSTMT": (*AlterUserMappingStmtMarshalAlias)(&node),
	})
}

func (node *AlterUserMappingStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterUserMappingStmt) Deparse() string {
	panic("Not Implemented")
}
