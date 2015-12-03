// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DeleteStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DeleteStmt")
	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	for _, subNode := range node.ReturningList {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.UsingClause {
		subNode.Fingerprint(ctx)
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx)
	}
}
