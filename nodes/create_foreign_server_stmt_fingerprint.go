// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateForeignServerStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateForeignServerStmt")

	if node.Fdwname != nil {
		ctx.WriteString("fdwname")
		ctx.WriteString(*node.Fdwname)
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Servername != nil {
		ctx.WriteString("servername")
		ctx.WriteString(*node.Servername)
	}

	if node.Servertype != nil {
		ctx.WriteString("servertype")
		ctx.WriteString(*node.Servertype)
	}

	if node.Version != nil {
		ctx.WriteString("version")
		ctx.WriteString(*node.Version)
	}
}
