// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node InsertStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "INSERT INTO")

	for _, subNode := range node.Cols {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	for _, subNode := range node.ReturningList {
		subNode.Fingerprint(ctx)
	}

	if node.SelectStmt != nil {
		node.SelectStmt.Fingerprint(ctx)
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx)
	}
}
