// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ForeignPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FOREIGNPATH")

	for _, subNode := range node.FdwPrivate {
		subNode.Fingerprint(ctx)
	}
}
