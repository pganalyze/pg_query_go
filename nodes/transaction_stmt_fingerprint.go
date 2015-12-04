// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node TransactionStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TRANSACTION")

	if node.Gid != nil {
		io.WriteString(ctx.hash, *node.Gid)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Kind)))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
