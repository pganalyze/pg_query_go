// Auto-generated - DO NOT EDIT

package pg_query

func (node AccessPriv) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AccessPriv")

	if len(node.Cols.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Cols.Fingerprint(&subCtx, "Cols")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("cols")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.PrivName != nil {
		ctx.WriteString("priv_name")
		ctx.WriteString(*node.PrivName)
	}
}
