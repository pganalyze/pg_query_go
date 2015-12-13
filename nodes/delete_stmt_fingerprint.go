// Auto-generated - DO NOT EDIT

package pg_query

func (node DeleteStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DeleteStmt")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	for _, subNode := range node.ReturningList {
		subNode.Fingerprint(ctx, "ReturningList")
	}

	for _, subNode := range node.UsingClause {
		subNode.Fingerprint(ctx, "UsingClause")
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
