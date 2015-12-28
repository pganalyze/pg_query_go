// Auto-generated - DO NOT EDIT

package pg_query

func (node ExplainStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ExplainStmt")
	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Query != nil {
		ctx.WriteString("query")
		node.Query.Fingerprint(ctx, "Query")
	}
}
