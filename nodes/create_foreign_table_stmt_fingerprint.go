// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateForeignTableStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateForeignTableStmt")
	node.Base.Fingerprint(ctx, "Base")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}
}
