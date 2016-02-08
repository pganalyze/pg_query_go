// Auto-generated - DO NOT EDIT

package pg_query

func (node DeleteStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DeleteStmt")

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

	if len(node.UsingClause.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.UsingClause.Fingerprint(&subCtx, "UsingClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("usingClause")
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
