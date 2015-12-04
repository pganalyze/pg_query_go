// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node BitmapOr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BITMAPOR")

	for _, subNode := range node.Bitmapplans {
		subNode.Fingerprint(ctx)
	}
}
