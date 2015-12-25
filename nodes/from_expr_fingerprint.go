// Auto-generated - DO NOT EDIT

package pg_query

func (node FromExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FromExpr")
	node.Fromlist.Fingerprint(ctx, "Fromlist")

	if node.Quals != nil {
		node.Quals.Fingerprint(ctx, "Quals")
	}
}
