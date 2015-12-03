// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RuleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RuleStmt")

	for _, subNode := range node.Actions {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}
}
