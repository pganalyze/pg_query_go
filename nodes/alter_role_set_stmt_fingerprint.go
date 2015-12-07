// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterRoleSetStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ALTERROLESETSTMT")

	if node.Database != nil {
		ctx.WriteString(*node.Database)
	}

	if node.Role != nil {
		ctx.WriteString(*node.Role)
	}

	if node.Setstmt != nil {
		node.Setstmt.Fingerprint(ctx)
	}
}
