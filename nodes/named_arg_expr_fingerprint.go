// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node NamedArgExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "NamedArgExpr")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
