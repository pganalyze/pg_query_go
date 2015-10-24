// Auto-generated - DO NOT EDIT

package pg_query

type TransactionStmt struct {
	Kind    TransactionStmtKind `json:"kind"`    /* see above */
	Options []Node              `json:"options"` /* for BEGIN/START and savepoint commands */
	Gid     *string             `json:"gid"`     /* for two-phase-commit related commands */
}
