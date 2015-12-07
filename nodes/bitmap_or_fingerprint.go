// Auto-generated - DO NOT EDIT

package pg_query

func (node BitmapOr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BITMAPOR")

	for _, subNode := range node.Bitmapplans {
		subNode.Fingerprint(ctx)
	}
}
