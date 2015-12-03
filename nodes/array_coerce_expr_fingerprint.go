// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ArrayCoerceExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ArrayCoerceExpr")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
