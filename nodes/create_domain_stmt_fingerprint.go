// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateDomainStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATEDOMAINSTMT")

	if node.CollClause != nil {
		node.CollClause.Fingerprint(ctx)
	}

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Domainname {
		subNode.Fingerprint(ctx)
	}

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}
}
