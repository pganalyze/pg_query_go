// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Append) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "APPEND")

	for _, subNode := range node.Appendplans {
		subNode.Fingerprint(ctx)
	}
}
