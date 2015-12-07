// Auto-generated - DO NOT EDIT

package pg_query

func (node PlaceHolderVar) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PLACEHOLDERVAR")

	if node.Phexpr != nil {
		node.Phexpr.Fingerprint(ctx)
	}
}
