// Auto-generated - DO NOT EDIT

package pg_query

func (node PrepareStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("PrepareStmt")

	if len(node.Argtypes.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Argtypes.Fingerprint(&subCtx, "Argtypes")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("argtypes")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
	// Intentionally ignoring node.Name for fingerprinting

	if node.Query != nil {
		subCtx := FingerprintSubContext{}
		node.Query.Fingerprint(&subCtx, "Query")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("query")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
