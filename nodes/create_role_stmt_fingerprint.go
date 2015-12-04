// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CreateRoleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATEROLESTMT")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Role != nil {
		io.WriteString(ctx.hash, *node.Role)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.StmtType)))
}
