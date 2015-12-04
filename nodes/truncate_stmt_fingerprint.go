// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node TruncateStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TRUNCATE")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Behavior)))

	for _, subNode := range node.Relations {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.RestartSeqs))
}
