// Auto-generated - DO NOT EDIT

package pg_query

func (node ColumnRef) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ColumnRef")

	for _, subNode := range node.Fields {
		subNode.Fingerprint(ctx, "Fields")
	}

	// Intentionally ignoring node.Location for fingerprinting
}
