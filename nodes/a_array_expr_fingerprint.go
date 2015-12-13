// Auto-generated - DO NOT EDIT

package pg_query

func (node A_ArrayExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("A_ArrayExpr")

	for _, subNode := range node.Elements {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}
