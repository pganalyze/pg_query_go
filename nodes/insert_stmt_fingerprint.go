// Auto-generated - DO NOT EDIT

package pg_query

func (node InsertStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("InsertStmt")

	if len(node.Cols.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Cols.Fingerprint(&subCtx, "Cols")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("cols")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.OnConflictClause != nil {
		subCtx := FingerprintSubContext{}
		node.OnConflictClause.Fingerprint(&subCtx, "OnConflictClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("onConflictClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Relation != nil {
		subCtx := FingerprintSubContext{}
		node.Relation.Fingerprint(&subCtx, "Relation")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("relation")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.ReturningList.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.ReturningList.Fingerprint(&subCtx, "ReturningList")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("returningList")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.SelectStmt != nil {
		subCtx := FingerprintSubContext{}
		node.SelectStmt.Fingerprint(&subCtx, "SelectStmt")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("selectStmt")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.WithClause != nil {
		subCtx := FingerprintSubContext{}
		node.WithClause.Fingerprint(&subCtx, "WithClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("withClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
