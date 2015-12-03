// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node UpdateStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "UpdateStmt")

	for _, subNode := range node.FromClause {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	for _, subNode := range node.ReturningList {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.TargetList {
		subNode.Fingerprint(ctx)
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx)
	}
}
