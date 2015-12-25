// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterFdwStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterFdwStmt")

	if node.Fdwname != nil {
		ctx.WriteString(*node.Fdwname)
	}

	node.FuncOptions.Fingerprint(ctx, "FuncOptions")
	node.Options.Fingerprint(ctx, "Options")
}
