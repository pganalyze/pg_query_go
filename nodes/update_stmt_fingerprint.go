// Auto-generated - DO NOT EDIT

package pg_query

func (node UpdateStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("UpdateStmt")

	for _, subNode := range node.FromClause {
		subNode.Fingerprint(ctx, "FromClause")
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	for _, subNode := range node.ReturningList {
		subNode.Fingerprint(ctx, "ReturningList")
	}

	for _, subNode := range node.TargetList {
		subNode.Fingerprint(ctx, "TargetList")
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
