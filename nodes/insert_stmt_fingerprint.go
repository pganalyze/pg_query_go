// Auto-generated - DO NOT EDIT

package pg_query

func (node InsertStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("InsertStmt")
	node.Cols.Fingerprint(ctx, "Cols")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	node.ReturningList.Fingerprint(ctx, "ReturningList")

	if node.SelectStmt != nil {
		node.SelectStmt.Fingerprint(ctx, "SelectStmt")
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
