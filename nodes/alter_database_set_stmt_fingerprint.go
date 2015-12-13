// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterDatabaseSetStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterDatabaseSetStmt")

	if node.Dbname != nil {
		ctx.WriteString(*node.Dbname)
	}

	if node.Setstmt != nil {
		node.Setstmt.Fingerprint(ctx, "Setstmt")
	}
}
