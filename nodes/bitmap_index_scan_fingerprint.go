// Auto-generated - DO NOT EDIT

package pg_query

func (node BitmapIndexScan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BITMAPINDEXSCAN")

	for _, subNode := range node.Indexqual {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexqualorig {
		subNode.Fingerprint(ctx)
	}
}
