// Auto-generated - DO NOT EDIT

package pg_query

func (node ExplainStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ExplainStmt")
	node.Options.Fingerprint(ctx, "Options")

	if node.Query != nil {
		node.Query.Fingerprint(ctx, "Query")
	}
}
