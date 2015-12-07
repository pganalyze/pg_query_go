// Auto-generated - DO NOT EDIT

package pg_query

func (node ArrayRef) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ARRAYREF")

	if node.Refassgnexpr != nil {
		node.Refassgnexpr.Fingerprint(ctx)
	}

	if node.Refexpr != nil {
		node.Refexpr.Fingerprint(ctx)
	}

	for _, subNode := range node.Reflowerindexpr {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Refupperindexpr {
		subNode.Fingerprint(ctx)
	}
}
