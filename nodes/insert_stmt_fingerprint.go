// Auto-generated - DO NOT EDIT

package pg_query

func (node InsertStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("InsertStmt")
	if len(node.Cols.Items) > 0 {
		ctx.WriteString("cols")
		node.Cols.Fingerprint(ctx, "Cols")
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if len(node.ReturningList.Items) > 0 {
		ctx.WriteString("returningList")
		node.ReturningList.Fingerprint(ctx, "ReturningList")
	}

	if node.SelectStmt != nil {
		ctx.WriteString("selectStmt")
		node.SelectStmt.Fingerprint(ctx, "SelectStmt")
	}

	if node.WithClause != nil {
		ctx.WriteString("withClause")
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
