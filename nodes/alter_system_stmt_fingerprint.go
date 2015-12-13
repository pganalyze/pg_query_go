// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterSystemStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterSystemStmt")

	if node.Setstmt != nil {
		node.Setstmt.Fingerprint(ctx)
	}
}
