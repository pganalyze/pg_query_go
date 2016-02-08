// Auto-generated - DO NOT EDIT

package pg_query

func (node ResTarget) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ResTarget")

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
	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil && parentFieldName != "TargetList" {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}

	if node.Val != nil {
		subCtx := FingerprintSubContext{}
		node.Val.Fingerprint(&subCtx, "Val")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("val")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
