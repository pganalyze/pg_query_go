// Auto-generated - DO NOT EDIT

package pg_query

func (node UpdateStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("UpdateStmt")

	if len(node.FromClause.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.FromClause.Fingerprint(&subCtx, "FromClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("fromClause")
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

	if len(node.TargetList.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.TargetList.Fingerprint(&subCtx, "TargetList")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("targetList")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.WhereClause != nil {
		subCtx := FingerprintSubContext{}
		node.WhereClause.Fingerprint(&subCtx, "WhereClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("whereClause")
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
