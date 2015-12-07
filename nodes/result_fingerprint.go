// Auto-generated - DO NOT EDIT

package pg_query

func (node Result) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RESULT")

	if node.Resconstantqual != nil {
		node.Resconstantqual.Fingerprint(ctx)
	}
}
