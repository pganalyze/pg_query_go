// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateFdwStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateFdwStmt")

	if node.Fdwname != nil {
		ctx.WriteString(*node.Fdwname)
	}

	node.FuncOptions.Fingerprint(ctx, "FuncOptions")
	node.Options.Fingerprint(ctx, "Options")
}
