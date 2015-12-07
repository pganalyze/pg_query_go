// Auto-generated - DO NOT EDIT

package pg_query

func (node CollateExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("COLLATE")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}
