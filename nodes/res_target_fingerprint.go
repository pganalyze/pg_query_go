// Auto-generated - DO NOT EDIT

package pg_query

func (node ResTarget) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RESTARGET")

	for _, subNode := range node.Indirection {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	if node.Val != nil {
		node.Val.Fingerprint(ctx)
	}
}
