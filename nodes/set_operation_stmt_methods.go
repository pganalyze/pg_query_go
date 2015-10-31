// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SetOperationStmt) MarshalJSON() ([]byte, error) {
	type SetOperationStmtMarshalAlias SetOperationStmt
	return json.Marshal(map[string]interface{}{
		"SETOPERATIONSTMT": (*SetOperationStmtMarshalAlias)(&node),
	})
}

func (node *SetOperationStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SetOperationStmt) Deparse() string {
	panic("Not Implemented")
}
