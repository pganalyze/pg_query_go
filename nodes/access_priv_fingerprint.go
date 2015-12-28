// Auto-generated - DO NOT EDIT

package pg_query

func (node AccessPriv) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AccessPriv")
	if len(node.Cols.Items) > 0 {
		ctx.WriteString("cols")
		node.Cols.Fingerprint(ctx, "Cols")
	}

	if node.PrivName != nil {
		ctx.WriteString("priv_name")
		ctx.WriteString(*node.PrivName)
	}
}
