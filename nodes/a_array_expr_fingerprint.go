// Auto-generated - DO NOT EDIT

package pg_query

func (node A_ArrayExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_ArrayExpr")

	for _, subNode := range node.Elements {
		subNode.Fingerprint(ctx, "Elements")
	}

	// Intentionally ignoring node.Location for fingerprinting
}
