// Auto-generated - DO NOT EDIT

package pg_query

func (node BitmapHeapPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BITMAPHEAPPATH")

	if node.Bitmapqual != nil {
		node.Bitmapqual.Fingerprint(ctx)
	}
}
