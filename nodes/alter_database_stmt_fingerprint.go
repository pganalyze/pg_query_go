// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterDatabaseStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterDatabaseStmt")

	if node.Dbname != nil {
		ctx.WriteString(*node.Dbname)
	}

	node.Options.Fingerprint(ctx, "Options")
}
