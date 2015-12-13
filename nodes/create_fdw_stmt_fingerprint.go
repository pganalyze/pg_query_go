// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateFdwStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateFdwStmt")

	if node.Fdwname != nil {
		ctx.WriteString(*node.Fdwname)
	}

	for _, subNode := range node.FuncOptions {
		subNode.Fingerprint(ctx, "FuncOptions")
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}
}
