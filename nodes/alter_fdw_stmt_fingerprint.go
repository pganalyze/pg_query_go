// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterFdwStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterFdwStmt")

	if node.Fdwname != nil {
		ctx.WriteString("fdwname")
		ctx.WriteString(*node.Fdwname)
	}

	if len(node.FuncOptions.Items) > 0 {
		ctx.WriteString("func_options")
		node.FuncOptions.Fingerprint(ctx, "FuncOptions")
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}
}
