// Auto-generated - DO NOT EDIT

package pg_query

func (node ExplainStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ExplainStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}
}
