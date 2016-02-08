// Auto-generated - DO NOT EDIT

package pg_query

func (node FromExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FromExpr")

	if len(node.Fromlist.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Fromlist.Fingerprint(&subCtx, "Fromlist")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("fromlist")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Quals != nil {
		subCtx := FingerprintSubContext{}
		node.Quals.Fingerprint(&subCtx, "Quals")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("quals")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
