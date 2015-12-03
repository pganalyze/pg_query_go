// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateDomainStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateDomainStmt")
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
