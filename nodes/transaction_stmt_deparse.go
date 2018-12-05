// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"fmt"
	"strings"
)

var transactionCmds = map[TransactionStmtKind]string{
	TRANS_STMT_BEGIN:             "BEGIN",
	TRANS_STMT_START:             "BEGIN",
	TRANS_STMT_COMMIT:            "COMMIT",
	TRANS_STMT_ROLLBACK:          "ROLLBACK",
	TRANS_STMT_SAVEPOINT:         "SAVEPOINT",
	TRANS_STMT_RELEASE:           "RELEASE",
	TRANS_STMT_ROLLBACK_TO:       "ROLLBACK",
	TRANS_STMT_PREPARE:           "PREPARE TRANSACTION",
	TRANS_STMT_COMMIT_PREPARED:   "COMMIT TRANSACTION",
	TRANS_STMT_ROLLBACK_PREPARED: "ROLLBACK TRANSACTION",
}

func (node TransactionStmt) Deparse(ctx DeparseContext) (string, error) {
	out := make([]string, 0)
	if kind, ok := transactionCmds[node.Kind]; !ok {
		return "", errors.New(fmt.Sprintf("couldn't deparse transaction kind: %d", node.Kind))
	} else {
		out = append(out, kind)
	}

	if node.Kind == TRANS_STMT_PREPARE ||
		node.Kind == TRANS_STMT_COMMIT_PREPARED ||
		node.Kind == TRANS_STMT_ROLLBACK_PREPARED {
		if node.Gid != nil {
			out = append(out, fmt.Sprintf("'%s'", *node.Gid))
		}
	} else {
		if node.Options.Items != nil && len(node.Options.Items) > 0 {

		}
	}
	return strings.Join(out, " "), nil
}
