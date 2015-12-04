// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node BitmapIndexScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BITMAPINDEXSCAN")

	for _, subNode := range node.Indexqual {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexqualorig {
		subNode.Fingerprint(ctx)
	}
}
