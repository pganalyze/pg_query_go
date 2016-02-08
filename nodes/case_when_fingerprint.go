// Auto-generated - DO NOT EDIT

package pg_query

func (node CaseWhen) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CaseWhen")

	if node.Expr != nil {
		subCtx := FingerprintSubContext{}
		node.Expr.Fingerprint(&subCtx, "Expr")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("expr")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
	// Intentionally ignoring node.Location for fingerprinting

	if node.Result != nil {
		subCtx := FingerprintSubContext{}
		node.Result.Fingerprint(&subCtx, "Result")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("result")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Xpr != nil {
		subCtx := FingerprintSubContext{}
		node.Xpr.Fingerprint(&subCtx, "Xpr")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("xpr")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
