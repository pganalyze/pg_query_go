// Auto-generated - DO NOT EDIT

package pg_query

func (node BitmapHeapScan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BITMAPHEAPSCAN")

	for _, subNode := range node.Bitmapqualorig {
		subNode.Fingerprint(ctx)
	}
}
