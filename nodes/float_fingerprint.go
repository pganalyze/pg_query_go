// Auto-generated - DO NOT EDIT

package pg_query

func (node Float) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Float")
	if len(node.Str) > 0 {
		ctx.WriteString("str")
		ctx.WriteString(node.Str)
	}
}
