// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FunctionParameter) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FunctionParameter")

	if node.ArgType != nil {
		node.ArgType.Fingerprint(ctx)
	}

	if node.Defexpr != nil {
		node.Defexpr.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Mode)))

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}
