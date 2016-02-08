// Auto-generated - DO NOT EDIT

package pg_query

func (node A_Indirection) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_Indirection")

	if node.Arg != nil {
		subCtx := FingerprintSubContext{}
		node.Arg.Fingerprint(&subCtx, "Arg")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("arg")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Indirection.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Indirection.Fingerprint(&subCtx, "Indirection")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("indirection")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
