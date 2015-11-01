// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type TransactionStmt struct {
	Kind    TransactionStmtKind `json:"kind"`    /* see above */
	Options []Node              `json:"options"` /* for BEGIN/START and savepoint commands */
	Gid     *string             `json:"gid"`     /* for two-phase-commit related commands */
}

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
