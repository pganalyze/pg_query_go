// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node DropStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DROP")

	for _, subNode := range node.Arguments {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Behavior)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Concurrent))
	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	for _, subNode := range node.Objects {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.RemoveType)))
}
