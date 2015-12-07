// Auto-generated - DO NOT EDIT

package pg_query

func (node BitmapOrPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BITMAPORPATH")

	for _, subNode := range node.Bitmapquals {
		subNode.Fingerprint(ctx)
	}
}
