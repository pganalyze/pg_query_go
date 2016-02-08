// Auto-generated - DO NOT EDIT

package pg_query

func (node DoStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DoStmt")

	if len(node.Args.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Args.Fingerprint(&subCtx, "Args")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("args")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
