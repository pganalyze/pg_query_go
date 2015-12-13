// Auto-generated - DO NOT EDIT

package pg_query

func (node ResTarget) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ResTarget")

	for _, subNode := range node.Indirection {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	// Intentionally ignoring node.Name for fingerprinting

	if node.Val != nil {
		node.Val.Fingerprint(ctx)
	}
}
