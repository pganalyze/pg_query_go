// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node InferenceElem) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("InferenceElem")

	if node.Expr != nil {
		subCtx := FingerprintSubContext{}
		node.Expr.Fingerprint(&subCtx, "Expr")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("expr")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Infercollid != 0 {
		ctx.WriteString("infercollid")
		ctx.WriteString(strconv.Itoa(int(node.Infercollid)))
	}

	if node.Inferopclass != 0 {
		ctx.WriteString("inferopclass")
		ctx.WriteString(strconv.Itoa(int(node.Inferopclass)))
	}

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
