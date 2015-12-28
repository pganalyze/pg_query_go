// Auto-generated - DO NOT EDIT

package pg_query

func (node ResTarget) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ResTarget")
	if len(node.Indirection.Items) > 0 {
		ctx.WriteString("indirection")
		node.Indirection.Fingerprint(ctx, "Indirection")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil && parentFieldName != "TargetList" {
		ctx.WriteString(*node.Name)
	}

	if node.Val != nil {
		ctx.WriteString("val")
		node.Val.Fingerprint(ctx, "Val")
	}
}
