// Auto-generated - DO NOT EDIT

package pg_query

func (node DeleteStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DeleteStmt")

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if len(node.ReturningList.Items) > 0 {
		ctx.WriteString("returningList")
		node.ReturningList.Fingerprint(ctx, "ReturningList")
	}

	if len(node.UsingClause.Items) > 0 {
		ctx.WriteString("usingClause")
		node.UsingClause.Fingerprint(ctx, "UsingClause")
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
