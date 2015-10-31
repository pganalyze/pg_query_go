// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DeallocateStmt) MarshalJSON() ([]byte, error) {
	type DeallocateStmtMarshalAlias DeallocateStmt
	return json.Marshal(map[string]interface{}{
		"DEALLOCATESTMT": (*DeallocateStmtMarshalAlias)(&node),
	})
}

func (node *DeallocateStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DeallocateStmt) Deparse() string {
	panic("Not Implemented")
}
