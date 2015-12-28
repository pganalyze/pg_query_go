// Auto-generated - DO NOT EDIT

package pg_query

func (node CreatedbStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreatedbStmt")

	if node.Dbname != nil {
		ctx.WriteString("dbname")
		ctx.WriteString(*node.Dbname)
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}
}
