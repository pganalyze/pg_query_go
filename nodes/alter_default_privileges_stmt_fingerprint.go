// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterDefaultPrivilegesStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterDefaultPrivilegesStmt")

	if node.Action != nil {
		ctx.WriteString("action")
		node.Action.Fingerprint(ctx, "Action")
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}
}
