// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FunctionParameter) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FunctionParameter")

	if node.ArgType != nil {
		node.ArgType.Fingerprint(ctx, "ArgType")
	}

	if node.Defexpr != nil {
		node.Defexpr.Fingerprint(ctx, "Defexpr")
	}

	ctx.WriteString(strconv.Itoa(int(node.Mode)))

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}
