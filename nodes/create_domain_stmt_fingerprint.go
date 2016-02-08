// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateDomainStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateDomainStmt")

	if node.CollClause != nil {
		subCtx := FingerprintSubContext{}
		node.CollClause.Fingerprint(&subCtx, "CollClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("collClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Constraints.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Constraints.Fingerprint(&subCtx, "Constraints")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("constraints")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Domainname.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Domainname.Fingerprint(&subCtx, "Domainname")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("domainname")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.TypeName != nil {
		subCtx := FingerprintSubContext{}
		node.TypeName.Fingerprint(&subCtx, "TypeName")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("typeName")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
