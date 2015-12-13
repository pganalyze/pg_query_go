// Auto-generated - DO NOT EDIT

package pg_query

func (node ExplainStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ExplainStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx, "Query")
	}
}
