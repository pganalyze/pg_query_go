// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node BitmapHeapScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BitmapHeapScan")

	for _, subNode := range node.Bitmapqualorig {
		subNode.Fingerprint(ctx)
	}
}
