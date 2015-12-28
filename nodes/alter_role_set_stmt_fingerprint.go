// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterRoleSetStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterRoleSetStmt")

	if node.Database != nil {
		ctx.WriteString("database")
		ctx.WriteString(*node.Database)
	}

	if node.Role != nil {
		ctx.WriteString("role")
		ctx.WriteString(*node.Role)
	}

	if node.Setstmt != nil {
		ctx.WriteString("setstmt")
		node.Setstmt.Fingerprint(ctx, "Setstmt")
	}
}
