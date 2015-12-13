// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterFdwStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterFdwStmt")

	if node.Fdwname != nil {
		ctx.WriteString(*node.Fdwname)
	}

	for _, subNode := range node.FuncOptions {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
