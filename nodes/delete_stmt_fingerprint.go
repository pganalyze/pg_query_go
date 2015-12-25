// Auto-generated - DO NOT EDIT

package pg_query

func (node DeleteStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DeleteStmt")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	node.ReturningList.Fingerprint(ctx, "ReturningList")
	node.UsingClause.Fingerprint(ctx, "UsingClause")

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
