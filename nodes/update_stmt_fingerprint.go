// Auto-generated - DO NOT EDIT

package pg_query

func (node UpdateStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("UpdateStmt")
	if len(node.FromClause.Items) > 0 {
		ctx.WriteString("fromClause")
		node.FromClause.Fingerprint(ctx, "FromClause")
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if len(node.ReturningList.Items) > 0 {
		ctx.WriteString("returningList")
		node.ReturningList.Fingerprint(ctx, "ReturningList")
	}

	if len(node.TargetList.Items) > 0 {
		ctx.WriteString("targetList")
		node.TargetList.Fingerprint(ctx, "TargetList")
	}

	if node.WhereClause != nil {
		ctx.WriteString("whereClause")
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}

	if node.WithClause != nil {
		ctx.WriteString("withClause")
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
