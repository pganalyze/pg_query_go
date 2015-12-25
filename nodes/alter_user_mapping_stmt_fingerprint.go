// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterUserMappingStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterUserMappingStmt")
	node.Options.Fingerprint(ctx, "Options")

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}

	if node.Username != nil {
		ctx.WriteString(*node.Username)
	}
}
