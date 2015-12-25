// Auto-generated - DO NOT EDIT

package pg_query

func (node DoStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DoStmt")
	node.Args.Fingerprint(ctx, "Args")
}
