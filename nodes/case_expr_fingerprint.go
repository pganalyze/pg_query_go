// Auto-generated - DO NOT EDIT

package pg_query

func (node CaseExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CASE")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	if node.Defresult != nil {
		node.Defresult.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}
