// Auto-generated - DO NOT EDIT

package pg_query

func (node CollateClause) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("COLLATECLAUSE")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	for _, subNode := range node.Collname {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}
