// Auto-generated - DO NOT EDIT

package pg_query

func (node CompositeTypeStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CompositeTypeStmt")

	for _, subNode := range node.Coldeflist {
		subNode.Fingerprint(ctx, "Coldeflist")
	}

	if node.Typevar != nil {
		node.Typevar.Fingerprint(ctx, "Typevar")
	}
}
