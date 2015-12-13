// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterExtensionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterExtensionStmt")

	if node.Extname != nil {
		ctx.WriteString(*node.Extname)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}
}
