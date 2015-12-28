// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterSystemStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterSystemStmt")

	if node.Setstmt != nil {
		ctx.WriteString("setstmt")
		node.Setstmt.Fingerprint(ctx, "Setstmt")
	}
}
