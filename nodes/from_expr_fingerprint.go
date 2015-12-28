// Auto-generated - DO NOT EDIT

package pg_query

func (node FromExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FromExpr")
	if len(node.Fromlist.Items) > 0 {
		ctx.WriteString("fromlist")
		node.Fromlist.Fingerprint(ctx, "Fromlist")
	}

	if node.Quals != nil {
		ctx.WriteString("quals")
		node.Quals.Fingerprint(ctx, "Quals")
	}
}
