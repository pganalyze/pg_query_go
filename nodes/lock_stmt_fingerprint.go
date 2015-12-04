// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node LockStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "LOCK")
	io.WriteString(ctx.hash, strconv.FormatBool(node.Nowait))

	for _, subNode := range node.Relations {
		subNode.Fingerprint(ctx)
	}
}
