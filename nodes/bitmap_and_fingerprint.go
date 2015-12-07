// Auto-generated - DO NOT EDIT

package pg_query

func (node BitmapAnd) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BITMAPAND")

	for _, subNode := range node.Bitmapplans {
		subNode.Fingerprint(ctx)
	}
}
