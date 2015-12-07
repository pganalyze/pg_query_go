// Auto-generated - DO NOT EDIT

package pg_query

func (node ExecuteStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("EXECUTESTMT")

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	for _, subNode := range node.Params {
		subNode.Fingerprint(ctx)
	}
}
