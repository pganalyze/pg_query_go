// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node BooleanTest) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("BooleanTest")

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

	if int(node.Booltesttype) != 0 {
		ctx.WriteString("booltesttype")
		ctx.WriteString(strconv.Itoa(int(node.Booltesttype)))
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
