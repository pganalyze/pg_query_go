// Auto-generated - DO NOT EDIT

package pg_query

func (node ColumnRef) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ColumnRef")
	if len(node.Fields.Items) > 0 {
		ctx.WriteString("fields")
		node.Fields.Fingerprint(ctx, "Fields")
	}

	// Intentionally ignoring node.Location for fingerprinting
}
