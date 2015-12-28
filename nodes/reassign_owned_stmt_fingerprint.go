// Auto-generated - DO NOT EDIT

package pg_query

func (node ReassignOwnedStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ReassignOwnedStmt")

	if node.Newrole != nil {
		ctx.WriteString("newrole")
		ctx.WriteString(*node.Newrole)
	}

	if len(node.Roles.Items) > 0 {
		ctx.WriteString("roles")
		node.Roles.Fingerprint(ctx, "Roles")
	}
}
