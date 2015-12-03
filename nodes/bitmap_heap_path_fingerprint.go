// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node BitmapHeapPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BitmapHeapPath")
	if node.Bitmapqual != nil {
		node.Bitmapqual.Fingerprint(ctx)
	}
}
