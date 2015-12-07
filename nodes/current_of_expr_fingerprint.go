// Auto-generated - DO NOT EDIT

package pg_query

func (node CurrentOfExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CURRENTOFEXPR")

	if node.CursorName != nil {
		ctx.WriteString(*node.CursorName)
	}
}
