// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateUserMappingStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateUserMappingStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}

	if node.Username != nil {
		ctx.WriteString(*node.Username)
	}
}
