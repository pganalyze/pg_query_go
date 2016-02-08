// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node A_Expr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_Expr")

	if int(node.Kind) != 0 {
		ctx.WriteString("kind")
		ctx.WriteString(strconv.Itoa(int(node.Kind)))
	}

	if node.Lexpr != nil {
		subCtx := FingerprintSubContext{}
		node.Lexpr.Fingerprint(&subCtx, "Lexpr")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("lexpr")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
	// Intentionally ignoring node.Location for fingerprinting

	if len(node.Name.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Name.Fingerprint(&subCtx, "Name")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("name")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Rexpr != nil {
		subCtx := FingerprintSubContext{}
		node.Rexpr.Fingerprint(&subCtx, "Rexpr")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("rexpr")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
