// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateForeignTableStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CreateForeignTableStmt")
	node.Base.Fingerprint(ctx)

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}
}
