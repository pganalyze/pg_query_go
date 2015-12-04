// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ForeignScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FOREIGNSCAN")

	for _, subNode := range node.FdwExprs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.FdwPrivate {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.FsSystemCol))
}
