// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateTableSpaceStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateTableSpaceStmt")
	// Intentionally ignoring node.Location for fingerprinting

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

	if node.Owner != nil {
		ctx.WriteString("owner")
		ctx.WriteString(*node.Owner)
	}

	if node.Tablespacename != nil {
		ctx.WriteString("tablespacename")
		ctx.WriteString(*node.Tablespacename)
	}
}
