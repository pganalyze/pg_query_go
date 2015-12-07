// Auto-generated - DO NOT EDIT

package pg_query

func (node WindowAgg) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("WINDOWAGG")

	if node.EndOffset != nil {
		node.EndOffset.Fingerprint(ctx)
	}

	if node.StartOffset != nil {
		node.StartOffset.Fingerprint(ctx)
	}
}
