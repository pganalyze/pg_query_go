// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node BitmapAnd) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BITMAPAND")

	for _, subNode := range node.Bitmapplans {
		subNode.Fingerprint(ctx)
	}
}
