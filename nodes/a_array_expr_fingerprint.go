// Auto-generated - DO NOT EDIT

package pg_query

func (node A_ArrayExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_ArrayExpr")
	if len(node.Elements.Items) > 0 {
		ctx.WriteString("elements")
		node.Elements.Fingerprint(ctx, "Elements")
	}

	// Intentionally ignoring node.Location for fingerprinting
}
