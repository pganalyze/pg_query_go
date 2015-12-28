// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateUserMappingStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateUserMappingStmt")
	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Servername != nil {
		ctx.WriteString("servername")
		ctx.WriteString(*node.Servername)
	}

	if node.Username != nil {
		ctx.WriteString("username")
		ctx.WriteString(*node.Username)
	}
}
