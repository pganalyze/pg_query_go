// Auto-generated - DO NOT EDIT

package pg_query

func (node CompositeTypeStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CompositeTypeStmt")
	if len(node.Coldeflist.Items) > 0 {
		ctx.WriteString("coldeflist")
		node.Coldeflist.Fingerprint(ctx, "Coldeflist")
	}

	if node.Typevar != nil {
		ctx.WriteString("typevar")
		node.Typevar.Fingerprint(ctx, "Typevar")
	}
}
