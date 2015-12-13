// Auto-generated - DO NOT EDIT

package pg_query

func (node ReassignOwnedStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ReassignOwnedStmt")

	if node.Newrole != nil {
		ctx.WriteString(*node.Newrole)
	}

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx, "Roles")
	}
}
