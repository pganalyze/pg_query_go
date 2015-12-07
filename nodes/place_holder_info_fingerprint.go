// Auto-generated - DO NOT EDIT

package pg_query

func (node PlaceHolderInfo) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PLACEHOLDERINFO")

	if node.PhVar != nil {
		node.PhVar.Fingerprint(ctx)
	}
}
