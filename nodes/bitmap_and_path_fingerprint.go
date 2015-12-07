// Auto-generated - DO NOT EDIT

package pg_query

func (node BitmapAndPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BITMAPANDPATH")

	for _, subNode := range node.Bitmapquals {
		subNode.Fingerprint(ctx)
	}
}
