// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateDomainStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateDomainStmt")

	if node.CollClause != nil {
		ctx.WriteString("collClause")
		node.CollClause.Fingerprint(ctx, "CollClause")
	}

	if len(node.Constraints.Items) > 0 {
		ctx.WriteString("constraints")
		node.Constraints.Fingerprint(ctx, "Constraints")
	}

	if len(node.Domainname.Items) > 0 {
		ctx.WriteString("domainname")
		node.Domainname.Fingerprint(ctx, "Domainname")
	}

	if node.TypeName != nil {
		ctx.WriteString("typeName")
		node.TypeName.Fingerprint(ctx, "TypeName")
	}
}
