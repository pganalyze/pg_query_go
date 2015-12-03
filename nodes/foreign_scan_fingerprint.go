// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ForeignScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ForeignScan")

	for _, subNode := range node.FdwExprs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.FdwPrivate {
		subNode.Fingerprint(ctx)
	}
}
