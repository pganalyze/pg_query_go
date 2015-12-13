// Auto-generated - DO NOT EDIT

package pg_query

func (node A_Indices) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("A_Indices")

	if node.Lidx != nil {
		node.Lidx.Fingerprint(ctx)
	}

	if node.Uidx != nil {
		node.Uidx.Fingerprint(ctx)
	}
}
