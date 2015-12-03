// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node BitmapAndPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BitmapAndPath")

	for _, subNode := range node.Bitmapquals {
		subNode.Fingerprint(ctx)
	}
}
