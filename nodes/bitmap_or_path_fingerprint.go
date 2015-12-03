// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node BitmapOrPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BitmapOrPath")

	for _, subNode := range node.Bitmapquals {
		subNode.Fingerprint(ctx)
	}
}
