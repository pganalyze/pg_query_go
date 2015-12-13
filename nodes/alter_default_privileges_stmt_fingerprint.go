// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterDefaultPrivilegesStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterDefaultPrivilegesStmt")

	if node.Action != nil {
		node.Action.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
