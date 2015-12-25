// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateDomainStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateDomainStmt")

	if node.CollClause != nil {
		node.CollClause.Fingerprint(ctx, "CollClause")
	}

	node.Constraints.Fingerprint(ctx, "Constraints")
	node.Domainname.Fingerprint(ctx, "Domainname")

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx, "TypeName")
	}
}
