// Auto-generated - DO NOT EDIT

package pg_query

func (node ColumnRef) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ColumnRef")
	node.Fields.Fingerprint(ctx, "Fields")
	// Intentionally ignoring node.Location for fingerprinting
}
