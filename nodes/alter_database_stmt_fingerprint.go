// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterDatabaseStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterDatabaseStmt")

	if node.Dbname != nil {
		ctx.WriteString(*node.Dbname)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}
}
