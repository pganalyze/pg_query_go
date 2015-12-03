// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node NestLoopParam) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "NestLoopParam")
	if node.Paramval != nil {
		node.Paramval.Fingerprint(ctx)
	}
}
