// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FunctionParameter) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FunctionParameter")

	if node.ArgType != nil {
		ctx.WriteString("argType")
		node.ArgType.Fingerprint(ctx, "ArgType")
	}

	if node.Defexpr != nil {
		ctx.WriteString("defexpr")
		node.Defexpr.Fingerprint(ctx, "Defexpr")
	}

	if int(node.Mode) != 0 {
		ctx.WriteString("mode")
		ctx.WriteString(strconv.Itoa(int(node.Mode)))
	}

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}
}
