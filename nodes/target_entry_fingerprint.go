// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node TargetEntry) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TargetEntry")
	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}
}
