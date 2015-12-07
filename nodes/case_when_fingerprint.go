// Auto-generated - DO NOT EDIT

package pg_query

func (node CaseWhen) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("WHEN")

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Result != nil {
		node.Result.Fingerprint(ctx)
	}
}
