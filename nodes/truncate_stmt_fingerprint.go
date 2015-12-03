// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node TruncateStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TruncateStmt")

	for _, subNode := range node.Relations {
		subNode.Fingerprint(ctx)
	}
}
