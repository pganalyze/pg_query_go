// Auto-generated - DO NOT EDIT

package pg_query

func (node UpdateStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("UpdateStmt")
	node.FromClause.Fingerprint(ctx, "FromClause")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	node.ReturningList.Fingerprint(ctx, "ReturningList")
	node.TargetList.Fingerprint(ctx, "TargetList")

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
