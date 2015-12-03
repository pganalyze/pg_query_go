// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node NestLoop) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "NestLoop")

	for _, subNode := range node.NestParams {
		subNode.Fingerprint(ctx)
	}
}
