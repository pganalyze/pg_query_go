// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node TransactionStmt) MarshalJSON() ([]byte, error) {
	type TransactionStmtMarshalAlias TransactionStmt
	return json.Marshal(map[string]interface{}{
		"TRANSACTION": (*TransactionStmtMarshalAlias)(&node),
	})
}

func (node *TransactionStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node TransactionStmt) Deparse() string {
	panic("Not Implemented")
}
