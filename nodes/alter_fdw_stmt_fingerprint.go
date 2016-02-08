// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterFdwStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterFdwStmt")

	if node.Fdwname != nil {
		ctx.WriteString("fdwname")
		ctx.WriteString(*node.Fdwname)
	}

	if len(node.FuncOptions.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.FuncOptions.Fingerprint(&subCtx, "FuncOptions")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("func_options")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Options.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Options.Fingerprint(&subCtx, "Options")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("options")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
