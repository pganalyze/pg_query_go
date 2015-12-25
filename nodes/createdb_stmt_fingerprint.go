// Auto-generated - DO NOT EDIT

package pg_query

func (node CreatedbStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreatedbStmt")

	if node.Dbname != nil {
		ctx.WriteString(*node.Dbname)
	}

	node.Options.Fingerprint(ctx, "Options")
}
