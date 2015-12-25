// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateUserMappingStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateUserMappingStmt")
	node.Options.Fingerprint(ctx, "Options")

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}

	if node.Username != nil {
		ctx.WriteString(*node.Username)
	}
}
