// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ConvertRowtypeExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ConvertRowtypeExpr")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
