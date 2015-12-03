// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node IndexStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "IndexStmt")

	for _, subNode := range node.ExcludeOpNames {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.IndexParams {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}
}
