// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node BoolExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("BoolExpr")

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

	if int(node.Boolop) != 0 {
		ctx.WriteString("boolop")
		ctx.WriteString(strconv.Itoa(int(node.Boolop)))
	}

	// Intentionally ignoring node.Location for fingerprinting

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
