// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterDefaultPrivilegesStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterDefaultPrivilegesStmt")

	if node.Action != nil {
		node.Action.Fingerprint(ctx, "Action")
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}
}
