// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ResultPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ResultPath")

	for _, subNode := range node.Quals {
		subNode.Fingerprint(ctx)
	}
}
