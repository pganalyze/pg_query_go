// Auto-generated - DO NOT EDIT

package pg_query

func (node CoalesceExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("COALESCE")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}
