// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node FunctionParameter) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FunctionParameter")
	if node.ArgType != nil {
		node.ArgType.Fingerprint(ctx)
	}

	if node.Defexpr != nil {
		node.Defexpr.Fingerprint(ctx)
	}
}
