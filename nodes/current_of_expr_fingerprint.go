// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CurrentOfExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CURRENTOFEXPR")

	if node.CursorName != nil {
		io.WriteString(ctx.hash, *node.CursorName)
	}
}
