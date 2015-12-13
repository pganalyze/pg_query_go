// Auto-generated - DO NOT EDIT

package pg_query

func (node ResTarget) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ResTarget")

	for _, subNode := range node.Indirection {
		subNode.Fingerprint(ctx, "Indirection")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil && parentFieldName != "TargetList" {
		ctx.WriteString(*node.Name)
	}

	if node.Val != nil {
		node.Val.Fingerprint(ctx, "Val")
	}
}
