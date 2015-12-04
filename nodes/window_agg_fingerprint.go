// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node WindowAgg) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WINDOWAGG")

	if node.EndOffset != nil {
		node.EndOffset.Fingerprint(ctx)
	}

	if node.StartOffset != nil {
		node.StartOffset.Fingerprint(ctx)
	}
}
